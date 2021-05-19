package server

import (
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

// injectUserToContext injects firebase user info to context
func injectUserToContext(db *db.Client, auth *auth.Client, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "" {
			bearerToken := strings.TrimPrefix(token, "Bearer ")
			firebaseToken, err := auth.VerifyIDTokenAndCheckRevoked(context.Background(), bearerToken)
			if err == nil {
				ctx := context.WithValue(r.Context(), request.CtxUserId, firebaseToken.UID)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	}
}

// NewRouterWithInject creates a new router with inject user middleware
func NewRouterWithInject(db *db.Client, auth *auth.Client, routers ...potato.Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = injectUserToContext(db, auth, route.HandlerFunc)
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
