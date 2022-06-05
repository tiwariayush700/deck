package handler

import (
	`context`

	`github.com/gin-gonic/gin`

	`deck/core/config`
	`deck/service`
)

type CardHandler struct {
	CardService   service.Card
	Configuration *config.Config
}

func NewCardHandler(cardService service.Card, configuration *config.Config) *CardHandler {
	return &CardHandler{
		CardService:   cardService,
		Configuration: configuration,
	}
}

func (h *CardHandler) RegisterRoutes(ctx context.Context, router *gin.Engine) {

	decksGroup := router.Group("/decks")
	{
		decksGroup.POST("", h.CreateDecks(ctx))
	}
}

func (h *CardHandler) CreateDecks(ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
