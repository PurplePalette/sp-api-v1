package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// NewFirebaseClient creates a new Firebase connection
func NewFirebaseClient() *firebase.App {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env, using os environment")
	}
	envCred := os.Getenv("FIREBASE_CREDENTIAL")
	opt := option.WithCredentialsJSON([]byte(envCred))
	envConf := os.Getenv("FIREBASE_CONFIG")
	var config *firebase.Config
	if err := json.Unmarshal([]byte(envConf), &config); err != nil {
		log.Fatal(err)
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}
	return app
}

// NewFirebaseFirestoreClient creates a new Firebase realtime database client
func NewFirebaseFirestoreClient(app *firebase.App) *firestore.Client {
	ctx := context.Background()
	firestore, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("error initializing database client:", err)
	}
	return firestore
}

// NewFirebaseAuthorizationClient creates a new Firebase authorization client
func NewFirebaseAuthorizationClient(app *firebase.App) *auth.Client {
	ctx := context.Background()
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatal("error initializing auth client:", err)
	}
	return auth
}
