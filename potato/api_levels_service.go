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
	"time"

	"cloud.google.com/go/firestore"
	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
	"gopkg.in/go-playground/validator.v9"
)

// LevelsAPIService is a service that implents the logic for the LevelsAPIServicer
// This service should implement the business logic for every endpoint for the LevelsAPI API.
// Include any external packages or services that will be required by this service.
type LevelsAPIService struct {
	firestore *firestore.Client
	cache     *CacheService
	validate  *validator.Validate
}

// NewLevelsAPIService creates a default api service
func NewLevelsAPIService(firestore *firestore.Client, cache *CacheService) LevelsAPIServicer {
	return &LevelsAPIService{firestore: firestore, cache: cache, validate: validator.New()}
}

// AddLevel - Add level
func (s *LevelsAPIService) AddLevel(ctx context.Context, levelName string, level Level) (ImplResponse, error) {
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
	userID, _ := request.GetUserID(ctx)
	level.UserID = userID
	level.Name = levelName
	currentTime := time.Now().Unix()
	level.CreatedTime = int32(currentTime)
	level.UpdatedTime = int32(currentTime)
	col := s.firestore.Collection("levels")
	// Add level to cache
	if err := s.cache.levels.Add(levelName, level); err != nil {
		return Response(http.StatusConflict, nil), nil
	}
	// Add level to firestore
	if _, err := col.Doc(levelName).Set(ctx, level); err != nil {
		log.Fatalln("Error posting level to firestore:", err)
		return Response(500, nil), nil
	}
	return Response(200, nil), nil
}

// EditLevel - Edit level
func (s *LevelsAPIService) EditLevel(ctx context.Context, levelName string, level Level) (ImplResponse, error) {
	if !request.IsValidName(levelName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	if err := s.validate.Struct(level); err != nil {
		return Response(http.StatusBadRequest, nil), nil
	}
	userID, _ := request.GetUserID(ctx)
	match, err := s.cache.levels.IsOwnerMatch(levelName, userID)
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	if !match {
		return Response(http.StatusForbidden, nil), nil
	}
	level.Name = levelName
	level.UpdatedTime = int32(time.Now().Unix())
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
func (s *LevelsAPIService) GetLevel(ctx context.Context, levelName string) (ImplResponse, error) {
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
func (s *LevelsAPIService) GetLevelList(ctx context.Context, localization string, page int32, keywords string) (ImplResponse, error) {
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
