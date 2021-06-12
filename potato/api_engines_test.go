package potato_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/server"
	"github.com/stretchr/testify/assert"
)

// CreateEnginesServer creates engine server for testing
func CreateEnginesServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}
	enginesAPIService := potato.NewEnginesAPIService(firestore, cache)
	enginesAPIController := potato.NewEnginesAPIController(enginesAPIService)
	router := server.NewRouterWithTestInject(auth, enginesAPIController)
	return httptest.NewServer(router)
}

func TestAddEngine(t *testing.T) {
	s := CreateEnginesServer()
	defer s.Close()
	ef := potato.Engine{
		Name:     "myEngine",
		Version:  1,
		Title:    "TestEngine",
		Subtitle: "only for test purpose",
		Author:   "Omado",
		Thumbnail: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			URL:  "",
		},
		Data: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			URL:  "",
		},
		Configuration: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			URL:  "",
		},
		Skin: potato.Skin{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			Texture:     potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		Background: potato.Background{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			Image:       potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		Effect: potato.Effect{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		Particle: potato.Particle{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			Texture:     potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		CreatedTime: 0,
		UpdatedTime: 0,
		UserID:      "",
	}
	efJson, _ := json.Marshal(ef)
	req := httptest.NewRequest(
		http.MethodPost,
		"/engines/myEngine",
		bytes.NewBuffer(efJson),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestEditEngine(t *testing.T) {
	s := CreateEnginesServer()
	defer s.Close()
	ef := potato.Engine{
		Name:     "myEngine",
		Version:  1,
		Title:    "TestEngine",
		Subtitle: "only for test purpose",
		Author:   "Omado",
		Thumbnail: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			URL:  "",
		},
		Data: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			URL:  "",
		},
		Configuration: potato.SonolusResourceLocator{
			Type: "",
			Hash: "",
			URL:  "",
		},
		Skin: potato.Skin{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			Texture:     potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		Background: potato.Background{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			Image:       potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		Effect: potato.Effect{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		Particle: potato.Particle{
			Name:        "",
			Version:     0,
			Title:       "",
			Subtitle:    "",
			Author:      "",
			Thumbnail:   potato.SonolusResourceLocator{},
			Data:        potato.SonolusResourceLocator{},
			Texture:     potato.SonolusResourceLocator{},
			CreatedTime: 0,
			UpdatedTime: 0,
			UserID:      "",
		},
		CreatedTime: 0,
		UpdatedTime: 0,
		UserID:      "",
	}
	efJson, _ := json.Marshal(ef)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/engines/myEngine",
		bytes.NewBuffer(efJson),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEngine(t *testing.T) {
	s := CreateEnginesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/engines/myEngine", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEngineList(t *testing.T) {
	s := CreateEnginesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/engines/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEngineListWithSpecifyKeyword(t *testing.T) {
	s := CreateEnginesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/engines/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "スキン2")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetEngineListWithSort(t *testing.T) {
	s := CreateEnginesServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/engines/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "sort:d order:d")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
