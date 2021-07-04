package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
)

// CacheService stores firestore data and provides add and remove methods to handle get request
type CacheService struct {
	// cacheInitService
	init *CacheInitService
	// backgrounds stores background
	backgrounds Cache
	// effects stores effect
	effects Cache
	// engines store engine
	engines Cache
	// levels stores level
	levels Cache
	// particles stores particle
	particles Cache
	// skins stores skin
	skins Cache
	// users stores user
	users Cache
	// tests stores testID map
	tests map[string]string
	// news stores fake levels
	news Cache
}

// NewCacheService creates a instance for stores firestore data
func NewCacheService(firestore *firestore.Client) *CacheService {
	s := NewCacheInitService(firestore)
	return &CacheService{
		init: s,
	}
}

// InitNews add news to cache instance (using static value for now)
func (s *CacheService) InitNews() {
	s.news.Data = make(map[string]interface{})
	_ = s.news.Add(
		"sweetPotatoWelcome",
		NewNews(
			"SweetPotatoサーバーへようこそ!",
			"もっと をタップして一覧を表示してください",
			"sweetPotatoWelcome",
			"PurplePalette DevTeam",
			410,
			"Cocoa",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/cocoa.jpg",
			"※この譜面は遊べません",
		),
	)
	_ = s.news.Add(
		"sweetPotatoUserWelcome",
		NewNews(
			"SweetPotatoユーザー個別サーバー",
			"Placeholder",
			"sweetPotatoUserWelcome",
			"PurplePalette DevTeam",
			1204,
			"Chino",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/chino.jpg",
			"※この譜面は遊べません",
		),
	)
	_ = s.news.Add(
		"sweetPotatoUserWelcome2",
		NewNews(
			"指定されたユーザーの投稿したデータのみ",
			"を遊ぶことができるサーバーです",
			"sweetPotatoUserWelcome2",
			"PurplePalette DevTeam",
			1204,
			"Chino",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/chino.jpg",
			"※この譜面は遊べません",
		),
	)
	_ = s.news.Add(
		"sweetPotatoTestWelcome",
		NewNews(
			"SweetPotatoテストサーバー",
			"Placeholder",
			"sweetPotatoTestWelcome",
			"PurplePalette DevTeam",
			214,
			"Rize",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/rize.jpg",
			"※この譜面は遊べません",
		),
	)
	_ = s.news.Add(
		"sweetPotatoTestWelcome2",
		NewNews(
			"まだ公開されていない投稿データのみ",
			"をテストプレイできるサーバーです",
			"sweetPotatoTestWelcome2",
			"PurplePalette DevTeam",
			214,
			"Rize",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/rize.jpg",
			"※この譜面は遊べません",
		),
	)
}

// InitCache initialize caches by getting whole firestore data (this method should only called once when start server)
func (s *CacheService) InitCache() error {
	backgrounds, err := s.init.LoadDatabaseFromFirebase("backgrounds")
	if err != nil {
		return err
	}
	s.backgrounds.Data = backgrounds
	effects, err := s.init.LoadDatabaseFromFirebase("effects")
	if err != nil {
		return errors.New("could not get effects from firestore")
	}
	s.effects.Data = effects
	engines, err := s.init.LoadDatabaseFromFirebase("engines")
	if err != nil {
		return errors.New("could not get engines from firestore")
	}
	s.engines.Data = engines
	levels, err := s.init.LoadDatabaseFromFirebase("levels")
	if err != nil {
		return errors.New("could not get levels from firestore")
	}
	s.levels.Data = levels
	particles, err := s.init.LoadDatabaseFromFirebase("particles")
	if err != nil {
		return errors.New("could not get particles from firestore")
	}
	s.particles.Data = particles
	skins, err := s.init.LoadDatabaseFromFirebase("skins")
	if err != nil {
		return errors.New("could not get skins from firestore")
	}
	s.skins.Data = skins
	users, err := s.init.LoadDatabaseFromFirebase("users")
	if err != nil {
		return errors.New("could not get user list from firestore")
	}
	s.users.Data = users
	s.tests = make(map[string]string)
	for _, user := range users {
		parsedUser := user.(User)
		s.tests[parsedUser.TestID] = parsedUser.UserID
	}
	s.InitNews()
	return nil
}

// GetUserIDFromTest gets the userID from the testID
// It returns the error if the testID wasn't valid
func (s *CacheService) GetUserIDFromTest(testID string) (string, error) {
	for _, user := range s.users.Data {
		parsedUser := user.(User)
		if parsedUser.TestID == testID {
			return parsedUser.UserID, nil
		}
	}
	return "", errors.New("could not find test")
}

// Add adds specified data to cache with using specified name as key.
// Data type must be background, effect, engine, level, particle, or skin.
func (s *CacheService) Add(name string, data interface{}) error {
	switch v := data.(type) {
	case Background:
		if err := s.backgrounds.Add(name, v); err != nil {
			return err
		}
	case Effect:
		if err := s.effects.Add(name, v); err != nil {
			return err
		}
	case Engine:
		if err := s.engines.Add(name, v); err != nil {
			return err
		}
	case Level:
		if err := s.levels.Add(name, v); err != nil {
			return err
		}
	case Particle:
		if err := s.particles.Add(name, v); err != nil {
			return err
		}
	case Skin:
		if err := s.skins.Add(name, v); err != nil {
			return err
		}
	case User:
		if err := s.users.Add(name, v); err != nil {
			return err
		}
	}
	return nil
}

// Set sets specified data to cache with using specified name as key.
// Data type must be background, effect, engine, level, particle, or skin.
func (s *CacheService) Set(name string, data interface{}) error {
	switch v := data.(type) {
	case Background:
		s.backgrounds.Set(name, v)
	case Effect:
		s.effects.Set(name, v)
	case Engine:
		s.engines.Set(name, v)
	case Level:
		s.levels.Set(name, v)
	case Particle:
		s.particles.Set(name, v)
	case Skin:
		s.skins.Set(name, v)
	case User:
		s.users.Set(name, v)
	default:
		return errors.New("unsupported data type")
	}
	return nil
}
