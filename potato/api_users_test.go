package potato_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/server"
	"github.com/stretchr/testify/assert"
)

// CreateUsersServer creates user filtered data endpoint server for testing
func CreateUsersServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	if err := potato.ReGenerateDatabase(firestore); err != nil {
		panic(err)
	}
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}
	usersApiService := potato.NewUsersApiService(firestore, cache)
	usersApiController := potato.NewUsersApiController(usersApiService)
	router := server.NewRouterWithTestInject(auth, usersApiController)
	return httptest.NewServer(router)
}

func TestGetUser(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUserList(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUserServerInfo(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/info", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersBackgrounds(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/backgrounds/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersEffects(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/effects/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersEngines(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/engines/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersLevels(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/levels/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	var levelsResponse potato.GetLevelListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &levelsResponse); err == nil {
		for _, level := range levelsResponse.Items {
			t.Log(level.Name)
		}
	}
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersParticles(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/particles/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersSkins(t *testing.T) {
	s := CreateUsersServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/users/YnaKWRpbanfyn1ge6FKQChqocyyn/skins/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
