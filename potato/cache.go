package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
)

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

func NewCacheService(firestore *firestore.Client) *CacheService {
	s := NewCacheInitService(firestore)
	return &CacheService{
		init: s,
	}
}

func (s *CacheService) InitNews() {
	s.news.Data = make(map[string]interface{})
	s.news.Add(
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
	s.news.Add(
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
	s.news.Add(
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
	s.news.Add(
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
	s.news.Add(
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

func (s *CacheService) InitCache() error {
	backgrounds, err := s.init.LoadDatabaseFromFirebase("backgrounds")
	if err != nil {
		return errors.New("could not get backgrounds from firestore")
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

func (c *CacheService) GetUserIDFromTest(testID string) (string, error) {
	userID, ok := c.tests[testID]
	if !ok {
		return "", errors.New("could not find test")
	}
	return userID, nil
}

func (c *CacheService) Add(name string, data interface{}) error {
	switch v := data.(type) {
	case Background:
		if err := c.backgrounds.Add(name, v); err != nil {
			return err
		}
	case Effect:
		if err := c.effects.Add(name, v); err != nil {
			return err
		}
	case Engine:
		if err := c.engines.Add(name, v); err != nil {
			return err
		}
	case Level:
		if err := c.levels.Add(name, v); err != nil {
			return err
		}
	case Particle:
		if err := c.particles.Add(name, v); err != nil {
			return err
		}
	case Skin:
		if err := c.skins.Add(name, v); err != nil {
			return err
		}
	}
	return nil
}

func (c *CacheService) Set(name string, data interface{}) error {
	switch v := data.(type) {
	case Background:
		c.backgrounds.Set(name, v)
	case Effect:
		c.effects.Set(name, v)
	case Engine:
		c.engines.Set(name, v)
	case Level:
		c.levels.Set(name, v)
	case Particle:
		c.particles.Set(name, v)
	case Skin:
		c.skins.Set(name, v)
	}
	return nil
}
