package server

import (
	"encoding/json"
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

// NewFirebaseDatabaseClient creates a new Firebase realtime database client
func NewFirebaseDatabaseClient(app *firebase.App) *db.Client {
	ctx := context.Background()
	db, err := app.Database(ctx)
	if err != nil {
		log.Fatal("error initializing database client:", err)
	}
	return db
}
