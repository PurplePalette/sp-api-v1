package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
)

type CacheService struct {
	// cacheInitService
	init *CacheInitService
	// userIdList stores user
	userList Cache
	// backgroundList stores background
	backgroundList Cache
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
	s.userList.Data = userList
	backgroundList, err := s.init.LoadBackgroundList()
	if err != nil {
		return errors.New("could not get background list from firestore")
	}
	s.backgroundList.Data = backgroundList
	return nil
}

// IsUserExist checks the user is exist and not removed
func (s *CacheService) IsUserExist(userId string) bool {
	return s.userList.IsExist(userId)
}

// IsBackgroundExist checks the user is exist and not removed
func (s *CacheService) IsBackgroundExist(bgName string) bool {
	return s.backgroundList.IsExist(bgName)
}
