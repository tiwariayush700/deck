package service

import (
	`context`

	`deck/core/api`
	`deck/model`
)

type Card interface {
	CreateDeck(ctx context.Context, request api.DeckRequest) (*model.Deck, error)
}