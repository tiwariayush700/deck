package repository

import (
	`context`

	`deck/core/database`
	`deck/model`
)

type CardRepository interface {
	Repository
	QueryCards(ctx context.Context, params map[string]interface{}, filter string, uow *UnitOfWork) ([]model.Card, database.Error)
	QueryInCards(ctx context.Context, v []string, uow *UnitOfWork) ([]model.Card, database.Error)
	GetDeckView(ctx context.Context, deckID string, uow *UnitOfWork) (*model.DeckView, database.Error)
}
