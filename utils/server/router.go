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
func injectUserToContext(auth *auth.Client, route potato.Route, sonolusVersion string) http.HandlerFunc {
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
			w.Header().Set("Sonolus-Version", sonolusVersion)
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

// indexHandler handles index page request.
// It returns redirect if INDEX_CONTENT was url.
// Otherwise, it return text directly.
// This endpoint won't support to specify html file, to prevent directory traversal.
func indexHandler(indexContent string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(indexContent, "http") {
			http.Redirect(w, r, indexContent, http.StatusMovedPermanently)
		} else {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(indexContent))
		}
	}
}

// NewRouterWithInject creates a new router with inject user middleware
func NewRouterWithInject(auth *auth.Client, routers ...potato.Router) *mux.Router {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env, using os environment")
	}
	indexContent := os.Getenv("INDEX_CONTENT")
	sonolusVersion := os.Getenv("SONOLUS_VERSION")
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = injectUserToContext(auth, route, sonolusVersion)
			handler = potato.Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	router.HandleFunc("/", indexHandler(indexContent))
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
