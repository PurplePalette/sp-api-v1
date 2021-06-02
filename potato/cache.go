package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
)

type CacheService struct {
	// cacheInitService
	init *CacheInitService
	// users stores user
	users Cache
	// backgrounds stores background
	backgrounds Cache
}

func NewCacheService(firestore *firestore.Client) *CacheService {
	s := NewCacheInitService(firestore)
	return &CacheService{
		init: s,
	}
}

func (s *CacheService) InitCache() error {
	userList, err := s.init.LoadUserList()
	if err != nil {
		return errors.New("could not get user list from firestore")
	}
	s.users.Data = userList
	backgrounds, err := s.init.LoadBackgroundList()
	if err != nil {
		return errors.New("could not get backgrounds from firestore")
	}
	s.backgrounds.Data = backgrounds
	return nil
}
