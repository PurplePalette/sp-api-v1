package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
)

// injectUserToContext injects firebase user info to context
func injectUserToContext(auth *auth.Client, route potato.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		authorized := false
		if token != "" {
			bearerToken := strings.TrimPrefix(token, "Bearer ")
			firebaseToken, err := auth.VerifyIDTokenAndCheckRevoked(context.Background(), bearerToken)
			if err == nil {
				ctx := context.WithValue(r.Context(), request.CtxUserID, firebaseToken.UID)
				r = r.WithContext(ctx)
				authorized = true
			}
		}
		if route.Method != "GET" && !authorized {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			route.HandlerFunc.ServeHTTP(w, r)
		}
	}
}

// injectTestUserToContext injects firebase user id to context (for test purpose)
func injectTestUserToContext(uid string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), request.CtxUserID, uid)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}

// NewRouterWithInject creates a new router with inject user middleware
func NewRouterWithInject(auth *auth.Client, routers ...potato.Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = injectUserToContext(auth, route)
			handler = potato.Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}

// NewRouterWithTestInject creates a new router with inject testUser middleware
func NewRouterWithTestInject(auth *auth.Client, routers ...potato.Router) *mux.Router {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env, using os environment")
	}
	uid := os.Getenv("TEST_UID")
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = injectTestUserToContext(uid, route.HandlerFunc)
			handler = potato.Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}
