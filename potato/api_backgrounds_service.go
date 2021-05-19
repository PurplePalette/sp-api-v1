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
	"errors"
	"log"
	"net/http"

	"firebase.google.com/go/db"
)

type FireUser struct {
	DateOfBirth string `json:"date_of_birth,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}

// BackgroundsApiService is a service that implents the logic for the BackgroundsApiServicer
// This service should implement the business logic for every endpoint for the BackgroundsApi API.
// Include any external packages or services that will be required by this service.
type BackgroundsApiService struct {
	db *db.Client
}

// NewBackgroundsApiService creates a default api service
func NewBackgroundsApiService(db *db.Client) BackgroundsApiServicer {
	return &BackgroundsApiService{db: db}
}

// AddBackground - Add background
func (s *BackgroundsApiService) AddBackground(ctx context.Context, backgroundName string, background Background) (ImplResponse, error) {
	// TODO - update AddBackground with the required logic for this service method.
	// Add api_backgrounds_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(409, {}) or use other options such as http.Ok ...
	//return Response(409, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddBackground method not implemented")
}

// EditBackground - Edit background
func (s *BackgroundsApiService) EditBackground(ctx context.Context, backgroundName string, background Background) (ImplResponse, error) {
	// TODO - update EditBackground with the required logic for this service method.
	// Add api_backgrounds_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditBackground method not implemented")
}

// GetBackground - Get background
func (s *BackgroundsApiService) GetBackground(ctx context.Context, backgroundName string) (ImplResponse, error) {
	// TODO - update GetBackground with the required logic for this service method.
	// Add api_backgrounds_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetBackgroundResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetBackgroundResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetBackground method not implemented")
}

// GetBackgroundList - Get background list
func (s *BackgroundsApiService) GetBackgroundList(ctx context.Context, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetBackgroundList with the required logic for this service method.
	// Add api_backgrounds_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetBackgroundListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetBackgroundListResponse{}), nil

	// Get a database reference to our blog.
	ref := s.db.NewRef("server/saving-data/fireblog")

	usersRef := ref.Child("users")
	err := usersRef.Set(ctx, map[string]*FireUser{
		"alanisawesome": {
			DateOfBirth: "June 23, 1912",
			FullName:    "Alan Turing",
		},
		"gracehop": {
			DateOfBirth: "December 9, 1906",
			FullName:    "Grace Hopper",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}

	return Response(http.StatusNotImplemented, nil), errors.New("GetBackgroundList method not implemented")
}
