package main

import (
	`context`
	`log`
	`time`

	`github.com/gin-gonic/gin`
	`gorm.io/gorm`

	`deck/core/config`
	`deck/core/database`
	`deck/handler`
	`deck/model`
	repositoryImpl `deck/repository/impl`
	`deck/service/impl`
	`deck/util`
)

type app struct {
	router        *gin.Engine
	configuration *config.Config
	pgDB          *gorm.DB
}

func newApp(router *gin.Engine, configuration *config.Config, pgDB *gorm.DB) *app {
	return &app{
		router:        router,
		configuration: configuration,
		pgDB:          pgDB,
	}
}

func (a *app) start(ctx context.Context) {

	//cardRepository
	cardRepository := repositoryImpl.NewCardRepositoryImpl()

	//cardService
	cardService := impl.NewCardServiceImpl(cardRepository, a.pgDB)

	//cardHandler
	cardHandler := handler.NewCardHandler(cardService, a.configuration)

	a.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//registerRoutes
	cardHandler.RegisterRoutes(ctx, a.router)

	log.Printf("Application loaded successfully on port : %s", a.configuration.Port)
	log.Fatal(a.router.Run(":" + a.configuration.Port))
}

func (a *app) autoMigrate() error {
	if err := a.pgDB.AutoMigrate(&model.Deck{}); err != nil {
		return err
	}

	if err := a.pgDB.AutoMigrate(&model.Card{}); err != nil {
		return err
	}

	if err := a.pgDB.AutoMigrate(&model.DeckCard{}); err != nil {
		return err
	}

	//migrate 52 cards
	for _, suit := range model.Suits {
		for rank := model.MinRank; rank <= model.MaxRank; rank++ {
			cardID := util.GetCode(suit, rank)
			code := util.GetCode(suit, rank)
			card := model.Card{
				BaseModel: database.BaseModel{
					ID:        cardID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Suit: suit,
				Rank: rank,
				Code: code,
			}

			if err := a.pgDB.FirstOrCreate(&card, "id = ?", card.ID).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
