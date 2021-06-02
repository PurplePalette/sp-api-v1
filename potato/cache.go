package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
)

type CacheService struct {
	// cacheInitService
	init *CacheInitService
	// userIdList stores user
	userList map[string]User
	// backgroundList stores background
	backgroundList map[string]Background
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
	s.userList = userList
	backgroundList, err := s.init.LoadBackgroundList()
	if err != nil {
		return errors.New("could not get background list from firestore")
	}
	s.backgroundList = backgroundList
	return nil
}

// IsUserExist checks the user is exist and not removed
func (s *CacheService) IsUserExist(userId string) bool {
	_, ok := s.userList[userId]
	return ok
}

// IsBackgroundExist checks the user is exist and not removed
func (s *CacheService) IsBackgroundExist(bgId string) bool {
	_, ok := s.backgroundList[bgId]
	return ok
}
