package potato

// News is fake level struct that has entire level struct
type News struct{ Level }

// NewNews creates a new fake level struct using specified texts
func NewNews(title string, subTitle string, upperLeftText string, upperRightText string, iconRating int, iconText string, iconURL string, description string) News {
	news := News{
		Level: Level{
			Name:    upperLeftText,
			Version: 1,
			Rating:  int32(iconRating),
			Engine: Engine{
				Name:          "",
				Version:       1,
				Title:         iconText,
				Subtitle:      "",
				Author:        "",
				Thumbnail:     SonolusResourceLocator{},
				Data:          SonolusResourceLocator{},
				Configuration: SonolusResourceLocator{},
				Skin:          Skin{},
				Background:    Background{},
				Effect:        Effect{},
				Particle:      Particle{},
				CreatedTime:   0,
				UpdatedTime:   0,
				UserID:        "",
			},
			UseSkin:       LevelUseSkin{},
			UseBackground: LevelUseBackground{},
			UseEffect:     LevelUseEffect{},
			UseParticle:   LevelUseParticle{},
			Title:         title,
			Artists:       subTitle,
			Author:        upperRightText,
			Cover: SonolusResourceLocator{
				Type: "LevelCover",
				Url:  iconURL,
			},
			Bgm:         SonolusResourceLocator{},
			Data:        SonolusResourceLocator{},
			Genre:       "",
			Public:      false,
			UserID:      "",
			Notes:       0,
			CreatedTime: 0,
			UpdatedTime: 0,
			Description: description,
		},
	}
	return news
}
