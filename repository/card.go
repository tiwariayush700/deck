package repository

import (
	`context`

	`deck/core/database`
	`deck/model`
)

type CardRepository interface {
	Repository
	QueryCards(ctx context.Context, params map[string]interface{}, filter string) ([]model.Card, database.Error)
	QueryInCards(ctx context.Context, v []string) ([]model.Card, database.Error)
	GetDeckView(ctx context.Context, deckID string, count int) (*model.DeckView, database.Error)
	DeleteCardsFromDeck(ctx context.Context, deckID string, cardIDs []string) database.Error
	CreateDeckCards(ctx context.Context, deck *model.Deck, cards []model.Card) database.Error
}
