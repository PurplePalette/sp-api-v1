/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openAPI-generator.tech)
 */

package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	potato "github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/server"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}

	BackgroundsAPIService := potato.NewBackgroundsAPIService(firestore, cache)
	BackgroundsAPIController := potato.NewBackgroundsAPIController(BackgroundsAPIService)

	EffectsAPIService := potato.NewEffectsAPIService(firestore, cache)
	EffectsAPIController := potato.NewEffectsAPIController(EffectsAPIService)

	EnginesAPIService := potato.NewEnginesAPIService(firestore, cache)
	EnginesAPIController := potato.NewEnginesAPIController(EnginesAPIService)

	InfoAPIService := potato.NewInfoAPIService(firestore, cache)
	InfoAPIController := potato.NewInfoAPIController(InfoAPIService)

	LevelsAPIService := potato.NewLevelsAPIService(firestore, cache)
	LevelsAPIController := potato.NewLevelsAPIController(LevelsAPIService)

	ParticlesAPIService := potato.NewParticlesAPIService(firestore, cache)
	ParticlesAPIController := potato.NewParticlesAPIController(ParticlesAPIService)

	SkinsAPIService := potato.NewSkinsAPIService(firestore, cache)
	SkinsAPIController := potato.NewSkinsAPIController(SkinsAPIService)

	TestsAPIService := potato.NewTestsAPIService(firestore, cache)
	TestsAPIController := potato.NewTestsAPIController(TestsAPIService)

	UsersAPIService := potato.NewUsersAPIService(firestore, cache)
	UsersAPIController := potato.NewUsersAPIController(UsersAPIService)

	listener := server.NewListener(firestore, cache)

	cols := []string{"backgrounds", "effects", "engines", "levels", "particles", "skins", "users"}
	for _, col := range cols {
		go listener.ListenFirestoreUpdate(col)
	}

	router := server.NewRouterWithInject(
		auth,
		BackgroundsAPIController,
		EffectsAPIController,
		EnginesAPIController,
		InfoAPIController,
		LevelsAPIController,
		ParticlesAPIController,
		SkinsAPIController,
		TestsAPIController,
		UsersAPIController,
	)

	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env, using os environment")
	}
	corsConf := os.Getenv("CORS_ORIGINS")
	allowedOrigins := strings.Split(corsConf, " ")
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	corsSupportedHandler := c.Handler(router)

	log.Printf("Server started!")
	log.Fatal(http.ListenAndServe(":8080", corsSupportedHandler))
}
