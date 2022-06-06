package service

import (
	`context`

	`deck/core/api`
	`deck/model`
)

type Card interface {
	CreateDeck(ctx context.Context, request api.DeckRequest) (*model.Deck, error)
	GetDeckView(ctx context.Context, count int, deckID string) (*model.DeckView, error)
	DrawCards(ctx context.Context, deckID string, cardIDs []string) error
}
