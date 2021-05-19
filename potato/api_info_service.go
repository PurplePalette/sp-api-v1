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
)

// InfoApiService is a service that implents the logic for the InfoApiServicer
// This service should implement the business logic for every endpoint for the InfoApi API.
// Include any external packages or services that will be required by this service.
type InfoApiService struct {
}

// NewInfoApiService creates a default api service
func NewInfoApiService() InfoApiServicer {
	return &InfoApiService{}
}

// EditInfo - Edit server info
func (s *InfoApiService) EditInfo(ctx context.Context, serverInfo ServerInfo) (ImplResponse, error) {
	// TODO - update EditInfo with the required logic for this service method.
	// Add api_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditInfo method not implemented")
}

// GetServerInfo - Get server info
func (s *InfoApiService) GetServerInfo(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetServerInfo with the required logic for this service method.
	// Add api_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ServerInfo{}) or use other options such as http.Ok ...
	//return Response(200, ServerInfo{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetServerInfo method not implemented")
}
