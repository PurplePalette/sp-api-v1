package potato

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// SetUserAuthorizationToHeader set requested user as you
func SetUserAuthorizationToHeader(req *http.Request) *http.Request {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env, using os environment")
	}
	token := os.Getenv("TEST_TOKEN")
	req.Header.Set("Authorization", "Bearer "+token)
	return req
}
