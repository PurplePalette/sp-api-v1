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

// CreateSkinsServer creates Skin server for testing
func CreateSkinsServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}
	SkinsApiService := potato.NewSkinsApiService(firestore, cache)
	SkinsApiController := potato.NewSkinsApiController(SkinsApiService)
	router := server.NewRouterWithTestInject(auth, SkinsApiController)
	return httptest.NewServer(router)
}

func TestAddSkin(t *testing.T) {
	s := CreateSkinsServer()
	defer s.Close()
	sk := potato.Skin{
		Name:        "mySkin",
		Version:     1,
		Title:       "HelloWorld",
		Subtitle:    "ex",
		Author:      "Omado",
		Thumbnail:   potato.SonolusResourceLocator{Type: "SkinThumbnail", Url: "https://example.com"},
		Data:        potato.SonolusResourceLocator{Type: "SkinData", Url: "https://example.com"},
		Texture:     potato.SonolusResourceLocator{Type: "SkinTexture", Url: "https://example.com"},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		UserId:      "YnaKWRpbanfyn1ge6FKQChqocyyn",
		Description: "スキンテスト",
	}
	skJson, _ := json.Marshal(sk)
	req := httptest.NewRequest(
		http.MethodPost,
		"/skins/mySkin",
		bytes.NewBuffer(skJson),
	)
	req = potato.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestEditSkin(t *testing.T) {
	s := CreateSkinsServer()
	defer s.Close()
	sk := potato.Skin{
		Name:        "mySkin",
		Version:     1,
		Title:       "HelloWorld2",
		Subtitle:    "ex",
		Author:      "Omado",
		Thumbnail:   potato.SonolusResourceLocator{Type: "SkinThumbnail", Url: "https://example.com"},
		Data:        potato.SonolusResourceLocator{Type: "SkinData", Url: "https://example.com"},
		Texture:     potato.SonolusResourceLocator{Type: "SkinTexture", Url: "https://example.com"},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		UserId:      "YnaKWRpbanfyn1ge6FKQChqocyyn",
		Description: "変更済みスキンテスト",
	}
	skJson, _ := json.Marshal(sk)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/skins/mySkin",
		bytes.NewBuffer(skJson),
	)
	req = potato.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetSkin(t *testing.T) {
	s := CreateSkinsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/skins/mySkin", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetSkinList(t *testing.T) {
	s := CreateSkinsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/skins/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetSkinListWithSpecifyKeyword(t *testing.T) {
	s := CreateSkinsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/skins/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "スキン2")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetSkinListWithSort(t *testing.T) {
	s := CreateSkinsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/skins/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "sort:d order:d")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
