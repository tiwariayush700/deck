package handler

import (
	`context`
	`net/http`

	`github.com/gin-gonic/gin`

	`deck/core/api`
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
		decksGroup.GET("/:id", h.GetDeck(ctx))
	}
}

func (h *CardHandler) CreateDecks(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		var deckRequest api.DeckRequest
		err := c.ShouldBind(&deckRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "Invalid request body"))
			return
		}

		deck, err := h.CardService.CreateDeck(ctx, deckRequest)
		if err != nil {
			if e, ok := err.(api.HTTPError); ok {
				if e.ErrorKey == api.ErrorCodeInvalidRequestPayload {
					c.JSON(http.StatusBadRequest, e)
					return
				}
				if e.ErrorKey == api.ErrorCodeResourceNotFound {
					c.JSON(http.StatusNotFound, e)
					return
				}
				if e.ErrorKey == api.ErrorCodeInternalError {
					c.JSON(http.StatusInternalServerError, e)
					return
				}
			}
			c.JSON(http.StatusBadRequest, api.NewHTTPError(api.ErrorCodeUnexpected, "Failed to create deck"))
			return
		}

		c.JSON(http.StatusOK, deck)
	}
}

func (h *CardHandler) GetDeck(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		deckID, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, api.NewHTTPError(api.ErrorCodeResourceNotFound, "No deck id passed"))
			return
		}

		deckView, err := h.CardService.GetDeckView(ctx, deckID)
		if err != nil {
			if e, ok := err.(api.HTTPError); ok {
				if e.ErrorKey == api.ErrorCodeInvalidRequestPayload {
					c.JSON(http.StatusBadRequest, e)
					return
				}
				if e.ErrorKey == api.ErrorCodeResourceNotFound {
					c.JSON(http.StatusNotFound, e)
					return
				}
				if e.ErrorKey == api.ErrorCodeInternalError {
					c.JSON(http.StatusInternalServerError, e)
					return
				}
			}
			c.JSON(http.StatusBadRequest, api.NewHTTPError(api.ErrorCodeUnexpected, "Failed to Fetch deck"))
			return
		}

		c.JSON(http.StatusOK, deckView)
	}
}
