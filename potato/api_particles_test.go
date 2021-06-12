package potato_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/server"
	"github.com/stretchr/testify/assert"
)

// CreateParticlesServer creates Particle server for testing
func CreateParticlesServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}
	ParticlesAPIService := potato.NewParticlesAPIService(firestore, cache)
	ParticlesAPIController := potato.NewParticlesAPIController(ParticlesAPIService)
	router := server.NewRouterWithTestInject(auth, ParticlesAPIController)
	return httptest.NewServer(router)
}

func TestAddParticle(t *testing.T) {
	s := CreateParticlesServer()
	defer s.Close()
	pt := potato.Particle{
		Name:        "myParticle",
		Version:     1,
		Title:       "HelloWorld",
		Subtitle:    "ex",
		Author:      "Omado",
		Thumbnail:   potato.SonolusResourceLocator{Type: "ParticleThumbnail", Url: "https://example.com"},
		Data:        potato.SonolusResourceLocator{Type: "ParticleData", Url: "https://example.com"},
		Texture:     potato.SonolusResourceLocator{Type: "ParticleTexture", Url: "https://example.com"},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		UserID:      "YnaKWRpbanfyn1ge6FKQChqocyyn",
		Description: "パーティクルテスト",
	}
	ptJson, _ := json.Marshal(pt)
	req := httptest.NewRequest(
		http.MethodPost,
		"/particles/myParticle",
		bytes.NewBuffer(ptJson),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestEditParticle(t *testing.T) {
	s := CreateParticlesServer()
	defer s.Close()
	pt := potato.Particle{
		Name:        "myParticle",
		Version:     1,
		Title:       "HelloWorld2",
		Subtitle:    "ex",
		Author:      "Omado",
		Thumbnail:   potato.SonolusResourceLocator{Type: "ParticleThumbnail", Url: "https://example.com"},
		Data:        potato.SonolusResourceLocator{Type: "ParticleData", Url: "https://example.com"},
		Texture:     potato.SonolusResourceLocator{Type: "ParticleTexture", Url: "https://example.com"},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		UserID:      "YnaKWRpbanfyn1ge6FKQChqocyyn",
		Description: "変更済みパーティクル",
	}
	ptJson, _ := json.Marshal(pt)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/particles/myParticle",
		bytes.NewBuffer(ptJson),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetParticle(t *testing.T) {
	s := CreateParticlesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/particles/myParticle", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetParticleList(t *testing.T) {
	s := CreateParticlesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/particles/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetParticleListWithSpecifyKeyword(t *testing.T) {
	s := CreateParticlesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/particles/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "スキン2")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetParticleListWithSort(t *testing.T) {
	s := CreateParticlesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/particles/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "sort:d order:d")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
