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

// CreateBackgroundsServer creates background server for testing
func CreateBackgroundsServer() *httptest.Server {
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
	BackgroundsAPIService := potato.NewBackgroundsAPIService(firestore, cache)
	BackgroundsAPIController := potato.NewBackgroundsAPIController(BackgroundsAPIService)
	router := server.NewRouterWithTestInject(auth, BackgroundsAPIController)
	return httptest.NewServer(router)
}

func TestAddBackground(t *testing.T) {
	s := CreateBackgroundsServer()
	defer s.Close()
	bg := potato.Background{
		Name:     "project-sekai",
		Version:  1,
		Title:    "プロセカスキン",
		Subtitle: "ほげほげ",
		Author:   "Anonymous",
		Thumbnail: potato.SonolusResourceLocator{
			Type: "BackgroundThumbnail",
			Url:  "",
		},
		Data: potato.SonolusResourceLocator{
			Type: "BackgroundData",
			Url:  "",
		},
		Image: potato.SonolusResourceLocator{
			Type: "BackgroundImage",
			Url:  "",
		},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
	}
	bgJson, _ := json.Marshal(bg)
	req := httptest.NewRequest(
		http.MethodPost,
		"/backgrounds/myBackground",
		bytes.NewBuffer(bgJson),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestEditBackground(t *testing.T) {
	s := CreateBackgroundsServer()
	defer s.Close()
	bg := potato.Background{
		Name:     "project-sekai",
		Version:  1,
		Title:    "プロセカスキン2",
		Subtitle: "ほげほげ",
		Author:   "Anonymous",
		Thumbnail: potato.SonolusResourceLocator{
			Type: "BackgroundThumbnail",
			Url:  "",
		},
		Data: potato.SonolusResourceLocator{
			Type: "BackgroundData",
			Url:  "",
		},
		Image: potato.SonolusResourceLocator{
			Type: "BackgroundImage",
			Url:  "",
		},
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
	}
	bgJson, _ := json.Marshal(bg)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/backgrounds/myBackground",
		bytes.NewBuffer(bgJson),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBackground(t *testing.T) {
	s := CreateBackgroundsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/backgrounds/myBackground", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBackgroundList(t *testing.T) {
	s := CreateBackgroundsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/backgrounds/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBackgroundListWithSpecifyKeyword(t *testing.T) {
	s := CreateBackgroundsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/backgrounds/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "スキン2")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBackgroundListWithSort(t *testing.T) {
	s := CreateBackgroundsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/backgrounds/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "sort:d order:d")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
