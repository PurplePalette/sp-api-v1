package server

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// NewFirebaseClient creates a new Firebase connection
func NewFirebaseClient() *firebase.App {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env, using os environment")
	}
	cred := os.Getenv("FIREBASE_CREDENTIAL")
	databaseUrl := os.Getenv("FIREBASE_DATABASE_URL")
	storageBucket := os.Getenv("FIREBASE_STORAGE_BUCKET")
	projectId := os.Getenv("FIREBASE_PROJECT_ID")
	serviceAccountId := os.Getenv("FIREBASE_SERVICE_ACCOUNT_ID")
	config := &firebase.Config{
		DatabaseURL:      databaseUrl,
		ProjectID:        projectId,
		ServiceAccountID: serviceAccountId,
		StorageBucket:    storageBucket,
	}
	opt := option.WithCredentialsJSON([]byte(cred))
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}
	return app
}

// NewFirebaseDatabaseClient creates a new Firebase realtime database client
func NewFirebaseDatabaseClient(app *firebase.App) *db.Client {
	ctx := context.Background()
	db, err := app.Database(ctx)
	if err != nil {
		log.Fatal("error initializing database client:", err)
	}
	return db
}
