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
	"github.com/PurplePalette/sonolus-uploader-core/utils/tests"
	"github.com/stretchr/testify/assert"
)

// CreateBackgroundsServer creates background server for testing
func CreateBackgroundsServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}

	BackgroundsApiService := potato.NewBackgroundsApiService(firestore, cache)
	BackgroundsApiController := potato.NewBackgroundsApiController(BackgroundsApiService)
	router := server.NewRouterWithTestInject(auth, BackgroundsApiController)
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
			Type: "BackgroundDat",
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
	req = tests.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
