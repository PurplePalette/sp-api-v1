package potato

type News struct{ Level }

func NewNews(title string, subTitle string, upperLeftText string, upperRightText string, iconRating int, iconText string, iconUrl string, description string) News {
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
				UserId:        "",
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
				Url:  iconUrl,
			},
			Bgm:         SonolusResourceLocator{},
			Data:        SonolusResourceLocator{},
			Genre:       "",
			Public:      false,
			UserId:      "",
			Notes:       0,
			CreatedTime: 0,
			UpdatedTime: 0,
			Description: description,
		},
	}
	return news
}
