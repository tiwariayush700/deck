package impl

import (
	`context`
	`log`

	`gorm.io/gorm`

	`deck/core/database`
	`deck/model`
	`deck/repository`
)

type cardRepositoryImpl struct {
	repositoryImpl
}

func (c *cardRepositoryImpl) GetDeckView(ctx context.Context, deckID string, uow *repository.UnitOfWork) (*model.DeckView, database.Error) {

	type response struct {
		DeckID    string     `gorm:"column:deck_id" json:"deck_id"`
		Shuffled  bool       `gorm:"column:deck_id" json:"shuffled"`
		Suit      model.Suit `gorm:"column:deck_id" json:"suit"`
		Rank      model.Rank `gorm:"column:rank" json:"rank"`
		Code      string     `gorm:"column:code" json:"code"`
		Remaining int        `gorm:"column:count" json:"remaining"`
	}
	views := make([]response, 0)
	err := uow.DB.Table("deck_cards").Select("deck_cards.deck_id, "+
		"decks.shuffled as shuffled, "+
		"decks.remaining as remaining, "+
		"cards.* as cards").
		Joins("INNER JOIN decks on deck_cards.deck_id = decks.id").
		Joins("INNER JOIN cards on deck_cards.card_id = cards.id").
		Where("deck_id = ?", deckID).
		Order("deck_cards.sequence_id asc").
		Find(&views).Error
	if err != nil {
		return nil, database.NewError(err)
	}

	var deckView model.DeckView
	cards := make([]model.CardView, 0)
	for _, view := range views {
		deckView.DeckID = view.DeckID
		deckView.Shuffled = view.Shuffled
		deckView.Remaining = view.Remaining
		cards = append(cards, model.CardView{Code: view.Code, Rank: view.Rank.String(), Suit: view.Suit.String()})
	}
	deckView.Cards = cards

	return &deckView, nil
}

func (c *cardRepositoryImpl) QueryInCards(ctx context.Context, v []string, uow *repository.UnitOfWork) ([]model.Card, database.Error) {
	cards := make([]model.Card, 0)
	err := uow.DB.Where("code IN ?", v).Find(&cards).Error
	if err != nil {
		log.Println("QueryInCards : err : ", err)
		return nil, database.NewError(err)
	}

	if len(cards) == 0 {
		return nil, database.NewError(gorm.ErrRecordNotFound)
	}

	return cards, nil
}

func (c *cardRepositoryImpl) QueryCards(ctx context.Context, params map[string]interface{}, filter string, uow *repository.UnitOfWork) ([]model.Card, database.Error) {
	cards := make([]model.Card, 0)
	err := uow.DB.Table("cards").
		Select("cards.*").
		Where(params).
		Where(filter).
		Find(&cards).Error

	if err != nil {
		log.Println("QueryCards : err : ", err)
		return nil, database.NewError(err)
	}

	if len(cards) == 0 {
		return nil, database.NewError(gorm.ErrRecordNotFound)
	}

	return cards, nil
}

func NewCardRepositoryImpl() repository.CardRepository {
	return &cardRepositoryImpl{}
}
