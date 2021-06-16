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

// BackgroundsAPIService is a service that implents the logic for the BackgroundsAPIServicer
// This service should implement the business logic for every endpoint for the BackgroundsAPI API.
// Include any external packages or services that will be required by this service.
type BackgroundsAPIService struct {
	firestore *firestore.Client
	cache     *CacheService
	validate  *validator.Validate
}

// NewBackgroundsAPIService creates a default api service
func NewBackgroundsAPIService(firestore *firestore.Client, cache *CacheService) BackgroundsAPIServicer {
	return &BackgroundsAPIService{firestore: firestore, cache: cache, validate: validator.New()}
}

// AddBackground - Add background
func (s *BackgroundsAPIService) AddBackground(ctx context.Context, backgroundName string, background Background) (ImplResponse, error) {
	if !request.IsValidName(backgroundName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	if err := s.validate.Struct(background); err != nil {
		return Response(http.StatusBadRequest, nil), nil
	}
	if s.cache.backgrounds.IsExist(backgroundName) {
		return Response(http.StatusConflict, nil), nil
	}
	// Force set parameter to valid
	userID, _ := request.GetUserID(ctx)
	background.UserID = userID
	background.Name = backgroundName
	nowTime := int32(time.Now().Unix())
	background.CreatedTime = nowTime
	background.UpdatedTime = nowTime
	col := s.firestore.Collection("backgrounds")
	// Add background to cache
	if err := s.cache.backgrounds.Add(backgroundName, background); err != nil {
		return Response(http.StatusConflict, nil), nil
	}
	// Add background to firestore
	if _, err := col.Doc(backgroundName).Set(ctx, background); err != nil {
		log.Println("Error posting background to firestore:", err)
		return Response(500, nil), nil
	}
	return Response(200, nil), nil
}

// EditBackground - Edit background
func (s *BackgroundsAPIService) EditBackground(ctx context.Context, backgroundName string, background Background) (ImplResponse, error) {
	if !request.IsValidName(backgroundName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	if err := s.validate.Struct(background); err != nil {
		return Response(http.StatusBadRequest, nil), nil
	}
	userID, _ := request.GetUserID(ctx)
	match, err := s.cache.backgrounds.IsOwnerMatch(backgroundName, userID)
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	if !match {
		return Response(http.StatusForbidden, nil), nil
	}
	// Update background data in firestore
	col := s.firestore.Collection("backgrounds")
	if _, err := col.Doc(backgroundName).Set(ctx, background); err != nil {
		log.Fatalln("Error posting background:", err)
		return Response(500, nil), nil
	}
	// Update background data in cache
	s.cache.backgrounds.Set(backgroundName, background)
	return Response(200, nil), nil
}

// GetBackground - Get background
func (s *BackgroundsAPIService) GetBackground(ctx context.Context, backgroundName string) (ImplResponse, error) {
	bg, err := s.cache.backgrounds.Get(backgroundName)
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	parsedBg := bg.(Background)
	resp := GetBackgroundResponse{
		Item:        parsedBg,
		Description: parsedBg.Description,
		Recommended: []Background{},
	}
	return Response(200, resp), nil
}

// GetBackgroundList - Get background list
func (s *BackgroundsAPIService) GetBackgroundList(ctx context.Context, localization string, page int32, keywords string) (ImplResponse, error) {
	query := request.ParseSearchQuery(keywords)
	pages := s.cache.backgrounds.Pages()
	items, err := s.cache.backgrounds.GetPage(page, query)
	if err != nil {
		log.Fatal(err)
		return Response(500, nil), nil
	}
	var backgrounds []Background
	err = json.Unmarshal(items, &backgrounds)
	if err != nil {
		return Response(500, nil), nil
	}
	resp := GetBackgroundListResponse{
		PageCount: pages,
		Items:     backgrounds,
	}
	return Response(200, resp), nil
}
