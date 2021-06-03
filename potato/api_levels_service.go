/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
	"gopkg.in/go-playground/validator.v9"
)

// LevelsApiService is a service that implents the logic for the LevelsApiServicer
// This service should implement the business logic for every endpoint for the LevelsApi API.
// Include any external packages or services that will be required by this service.
type LevelsApiService struct {
	firestore *firestore.Client
	cache     *CacheService
	validate  *validator.Validate
}

// NewLevelsApiService creates a default api service
func NewLevelsApiService(firestore *firestore.Client, cache *CacheService) LevelsApiServicer {
	return &LevelsApiService{firestore: firestore, cache: cache, validate: validator.New()}
}

// AddLevel - Add level
func (s *LevelsApiService) AddLevel(ctx context.Context, levelName string, level Level) (ImplResponse, error) {
	if !request.IsLoggedIn(ctx) {
		return Response(http.StatusUnauthorized, nil), nil
	}
	if !request.IsValidName(levelName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	if err := s.validate.Struct(level); err != nil {
		return Response(http.StatusBadRequest, err.Error()), nil
	}
	if s.cache.levels.IsExist(levelName) {
		return Response(http.StatusConflict, nil), nil
	}
	// Force set parameter to valid
	userId, _ := request.GetUserId(ctx)
	level.UserId = userId
	level.Name = levelName
	col := s.firestore.Collection("levels")
	// Add level to firestore
	if _, err := col.Doc(levelName).Set(ctx, level); err != nil {
		log.Fatalln("Error posting level:", err)
		return Response(500, nil), nil
	}
	// Add level to cache
	s.cache.levels.Add(levelName, level)
	return Response(200, nil), nil
}

// EditLevel - Edit level
func (s *LevelsApiService) EditLevel(ctx context.Context, levelName string, level Level) (ImplResponse, error) {
	if !request.IsLoggedIn(ctx) {
		return Response(http.StatusUnauthorized, nil), nil
	}
	if !request.IsValidName(levelName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	if err := s.validate.Struct(level); err != nil {
		return Response(http.StatusBadRequest, nil), nil
	}
	userId, _ := request.GetUserId(ctx)
	match, err := s.cache.levels.IsOwnerMatch(levelName, userId)
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	if !match {
		return Response(http.StatusForbidden, nil), nil
	}
	level.Name = levelName
	// Update level data in firestore
	col := s.firestore.Collection("levels")
	if _, err := col.Doc(levelName).Set(ctx, level); err != nil {
		log.Fatalln("Error posting level:", err)
		return Response(500, nil), nil
	}
	// Update level data in cache
	s.cache.levels.Set(levelName, level)
	return Response(200, nil), nil
}

// GetLevel - Get level
func (s *LevelsApiService) GetLevel(ctx context.Context, levelName string) (ImplResponse, error) {
	rawNs, newsNotExistErr := s.cache.news.Get(levelName)
	rawLv, levelNotExistErr := s.cache.levels.Get(levelName)
	if newsNotExistErr != nil && levelNotExistErr != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	var lv Level
	if newsNotExistErr == nil {
		ns := rawNs.(News)
		lv = ns.Level
	} else {
		lv = rawLv.(Level)
	}
	resp := GetLevelResponse{
		Item:        lv,
		Description: lv.Description,
		Recommended: []Level{},
	}
	return Response(200, resp), nil
}

// GetLevelList - Get level list
func (s *LevelsApiService) GetLevelList(ctx context.Context, localization string, page int32, keywords string) (ImplResponse, error) {
	query := request.ParseSearchQuery(keywords)
	pages := s.cache.levels.Pages()
	items, err := s.cache.levels.GetPage(page, query)
	if err != nil {
		log.Fatal(err)
		return Response(500, nil), nil
	}
	var levels []Level
	err = json.Unmarshal(items, &levels)
	if err != nil {
		return Response(500, nil), nil
	}
	resp := GetLevelListResponse{
		PageCount: pages,
		Items:     levels,
	}
	return Response(200, resp), nil
}
