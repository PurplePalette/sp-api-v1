/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	potato "github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/server"
)

func main() {
	firebase := server.NewFirebaseClient()
	db := server.NewFirebaseDatabaseClient(firebase)
	log.Printf("Server started")

	BackgroundsApiService := potato.NewBackgroundsApiService(db)
	BackgroundsApiController := potato.NewBackgroundsApiController(BackgroundsApiService)

	EffectsApiService := potato.NewEffectsApiService(db)
	EffectsApiController := potato.NewEffectsApiController(EffectsApiService)

	EnginesApiService := potato.NewEnginesApiService(db)
	EnginesApiController := potato.NewEnginesApiController(EnginesApiService)

	InfoApiService := potato.NewInfoApiService(db)
	InfoApiController := potato.NewInfoApiController(InfoApiService)

	LevelsApiService := potato.NewLevelsApiService(db)
	LevelsApiController := potato.NewLevelsApiController(LevelsApiService)

	ParticlesApiService := potato.NewParticlesApiService(db)
	ParticlesApiController := potato.NewParticlesApiController(ParticlesApiService)

	SkinsApiService := potato.NewSkinsApiService(db)
	SkinsApiController := potato.NewSkinsApiController(SkinsApiService)

	TestsApiService := potato.NewTestsApiService(db)
	TestsApiController := potato.NewTestsApiController(TestsApiService)

	UsersApiService := potato.NewUsersApiService(db)
	UsersApiController := potato.NewUsersApiController(UsersApiService)

	router := potato.NewRouter(BackgroundsApiController, EffectsApiController, EnginesApiController, InfoApiController, LevelsApiController, ParticlesApiController, SkinsApiController, TestsApiController, UsersApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
