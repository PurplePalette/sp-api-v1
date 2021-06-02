package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
)

type CacheInitService struct {
	firestore *firestore.Client
}

func NewCacheInitService(firestore *firestore.Client) *CacheInitService {
	return &CacheInitService{firestore: firestore}
}

// LoadDatabaseFromFirebase gets the entire database from Firebase
func (s *CacheInitService) LoadDatabaseFromFirebase(colName string, response *map[string]interface{}) error {
	// TODO: 他の取得処理をここにまとめる
	return nil
}

// LoadUserList gets the list of users from firebase for caching user status
func (s *CacheInitService) LoadUserList() (map[string]User, error) {
	userList := map[string]User{}
	col := s.firestore.Collection("users")
	docs, err := col.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, errors.New("could not get user collection from firestore")
	}
	for _, doc := range docs {
		var user User
		if err := doc.DataTo(&user); err != nil {
			return nil, errors.New("could not parse doc to user struct")
		}
		userList[doc.Ref.ID] = user
	}
	return userList, nil
}

// LoadBackgroundList gets the list of backgrounds from firebase for caching status
func (s *CacheInitService) LoadBackgroundList() (map[string]Background, error) {
	bgList := map[string]Background{}
	col := s.firestore.Collection("backgrounds")
	docs, err := col.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, errors.New("could not get background collection from firestore")
	}
	for _, doc := range docs {
		var bg Background
		if err := doc.DataTo(&bg); err != nil {
			return nil, errors.New("could not parse doc to background struct")
		}
		bgList[doc.Ref.ID] = bg
	}
	return bgList, nil
}
