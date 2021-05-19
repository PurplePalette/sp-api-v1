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
	"net/http"

	"firebase.google.com/go/db"
)

// UsersApiService is a service that implents the logic for the UsersApiServicer
// This service should implement the business logic for every endpoint for the UsersApi API.
// Include any external packages or services that will be required by this service.
type UsersApiService struct {
	db *db.Client
}

// NewUsersApiService creates a default api service
func NewUsersApiService(db *db.Client) UsersApiServicer {
	return &UsersApiService{db: db}
}

// EditUser - Edit user
func (s *UsersApiService) EditUser(ctx context.Context, userId string, user User) (ImplResponse, error) {
	// TODO - update EditUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

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

	return Response(http.StatusNotImplemented, nil), errors.New("EditUser method not implemented")
}

// GetUser - Get user
func (s *UsersApiService) GetUser(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update GetUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	//return Response(200, User{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUser method not implemented")
}

// GetUserList - Get user list
func (s *UsersApiService) GetUserList(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetUserList with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetUserListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetUserListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserList method not implemented")
}

// GetUserServerInfo - Get user server info
func (s *UsersApiService) GetUserServerInfo(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update GetUserServerInfo with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ServerInfo{}) or use other options such as http.Ok ...
	//return Response(200, ServerInfo{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserServerInfo method not implemented")
}

// GetUsersBackgrounds - Get backgrounds for test
func (s *UsersApiService) GetUsersBackgrounds(ctx context.Context, userId string, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetUsersBackgrounds with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetBackgroundListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetBackgroundListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersBackgrounds method not implemented")
}

// GetUsersEffects - Get effects for test
func (s *UsersApiService) GetUsersEffects(ctx context.Context, userId string, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetUsersEffects with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetEffectListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetEffectListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersEffects method not implemented")
}

// GetUsersEngines - Get engines for test
func (s *UsersApiService) GetUsersEngines(ctx context.Context, userId string, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetUsersEngines with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetEngineListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetEngineListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersEngines method not implemented")
}

// GetUsersLevels - Get levels for test
func (s *UsersApiService) GetUsersLevels(ctx context.Context, userId string, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetUsersLevels with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetLevelListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetLevelListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersLevels method not implemented")
}

// GetUsersParticles - Get particles for test
func (s *UsersApiService) GetUsersParticles(ctx context.Context, userId string, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetUsersParticles with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetParticleListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetParticleListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersParticles method not implemented")
}

// GetUsersSkins - Get skins for test
func (s *UsersApiService) GetUsersSkins(ctx context.Context, userId string, localization string, page int32, keywords string) (ImplResponse, error) {
	// TODO - update GetUsersSkins with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetSkinListResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetSkinListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersSkins method not implemented")
}
