package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
)

// CacheInitService is service for getting data from firestore
type CacheInitService struct {
	firestore *firestore.Client
}

// NewCacheInitService creates a new instance for getting data from firestore
func NewCacheInitService(firestore *firestore.Client) *CacheInitService {
	return &CacheInitService{firestore: firestore}
}

// LoadDatabaseFromFirebase gets the entire database from Firebase
func (s *CacheInitService) LoadDatabaseFromFirebase(colName string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	col := s.firestore.Collection(colName)
	docs, err := col.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, errors.New("could not get user collection from firestore")
	}
	for _, doc := range docs {
		switch colName {
		case "backgrounds":
			var bg Background
			if err := doc.DataTo(&bg); err != nil {
				return nil, errors.New("could not parse doc to background struct")
			}
			data[bg.Name] = bg
		case "effects":
			var ef Effect
			if err := doc.DataTo(&ef); err != nil {
				return nil, errors.New("could not parse doc to effect struct")
			}
			data[ef.Name] = ef
		case "engines":
			var eg Engine
			if err := doc.DataTo(&eg); err != nil {
				return nil, errors.New("could not parse doc to engine struct")
			}
			data[eg.Name] = eg
		case "levels":
			var lv Level
			if err := doc.DataTo(&lv); err != nil {
				return nil, errors.New("could not parse doc to level struct")
			}
			data[lv.Name] = lv
		case "particles":
			var pt Particle
			if err := doc.DataTo(&pt); err != nil {
				return nil, errors.New("could not parse doc to particle struct")
			}
			data[pt.Name] = pt
		case "skins":
			var sk Skin
			if err := doc.DataTo(&sk); err != nil {
				return nil, errors.New("could not parse doc to skin struct")
			}
			data[sk.Name] = sk
		case "users":
			var ur User
			if err := doc.DataTo(&ur); err != nil {
				return nil, errors.New("could not parse doc to user struct")
			}
			data[ur.UserID] = ur
		}
	}
	return data, nil
}
