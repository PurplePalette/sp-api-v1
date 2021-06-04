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

// CreateEffectsServer creates Effect server for testing
func CreateEffectsServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}
	EffectsApiService := potato.NewEffectsApiService(firestore, cache)
	EffectsApiController := potato.NewEffectsApiController(EffectsApiService)
	router := server.NewRouterWithTestInject(auth, EffectsApiController)
	return httptest.NewServer(router)
}

func TestAddEffect(t *testing.T) {
	s := CreateEffectsServer()
	defer s.Close()
	ef := potato.Effect{
		Name:     "myEffect",
		Version:  1,
		Title:    "Hello world",
		Subtitle: "ex",
		Author:   "Omado",
		Thumbnail: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			Url:  "",
		},
		Data: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			Url:  "",
		},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		UserId:      "kfcn",
	}
	efJson, _ := json.Marshal(ef)
	req := httptest.NewRequest(
		http.MethodPost,
		"/effects/myEffect",
		bytes.NewBuffer(efJson),
	)
	req = potato.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestEditEffect(t *testing.T) {
	s := CreateEffectsServer()
	defer s.Close()
	ef := potato.Effect{
		Name:     "myEffect",
		Version:  1,
		Title:    "Hello world",
		Subtitle: "ex",
		Author:   "Omado",
		Thumbnail: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			Url:  "",
		},
		Data: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			Url:  "",
		},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		UserId:      "kfcn",
	}
	efJson, _ := json.Marshal(ef)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/effects/myEffect",
		bytes.NewBuffer(efJson),
	)
	req = potato.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEffect(t *testing.T) {
	s := CreateEffectsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/effects/myEffect", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEffectList(t *testing.T) {
	s := CreateEffectsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/effects/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEffectListWithSpecifyKeyword(t *testing.T) {
	s := CreateEffectsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/effects/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "スキン2")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEffectListWithSort(t *testing.T) {
	s := CreateEffectsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/effects/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "sort:d order:d")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
