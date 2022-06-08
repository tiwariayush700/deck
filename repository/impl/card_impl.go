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
	db *gorm.DB
}

func (c *cardRepositoryImpl) CreateDeckCards(ctx context.Context, deck *model.Deck, cards []model.Card) database.Error {

	uow := repository.NewUnitOfWork(c.db, false)
	defer uow.Complete()

	dbErr := uow.DB.Create(deck).Error
	if dbErr != nil {
		log.Printf("Failed to Create deck : err %v", dbErr)
		return database.NewError(dbErr)
	}

	deckCards := make([]model.DeckCard, 0)
	for _, card := range cards {
		deckCards = append(deckCards, *model.NewDeckCard(deck.ID, card.ID))
	}

	dbErr = uow.DB.Create(&deckCards).Error
	if dbErr != nil {
		log.Printf("Failed to Create deckCards : err %v", dbErr)
		return database.NewError(dbErr)
	}

	err := uow.Commit()
	if err != nil {
		log.Printf("Failed to commit unit of work : err %v", err)
		return database.NewError(err)
	}

	return nil
}

func (c *cardRepositoryImpl) DeleteCardsFromDeck(ctx context.Context, deckID string, cardIDs []string) database.Error {

	uow := repository.NewUnitOfWork(c.db, false)
	defer uow.Complete()

	log.Printf("CardIds => %v", cardIDs)
	updatedCount := uow.DB.Where(" deck_id = ? AND card_id IN ? ", deckID, cardIDs).Delete(&model.DeckCard{}).RowsAffected
	if updatedCount == 0 {
		return database.NewError(gorm.ErrRecordNotFound)
	}

	dbErr := uow.DB.Model(&model.Deck{}).Where("id = ? ", deckID).Update("remaining", gorm.Expr("remaining - ?", updatedCount)).Error
	if dbErr != nil {
		log.Printf("Failed to Update remaining count : err %v", dbErr)
		return database.NewError(dbErr)
	}

	err := uow.Commit()
	if err != nil {
		log.Printf("Failed to commit unit of work : err %v", err)
		return database.NewError(err)
	}

	return nil
}

//todo add count param
func (c *cardRepositoryImpl) GetDeckView(ctx context.Context, deckID string, count int) (*model.DeckView, database.Error) {

	type response struct {
		DeckID    string     `gorm:"column:deck_id" json:"deck_id"`
		Shuffled  bool       `gorm:"column:deck_id" json:"shuffled"`
		Suit      model.Suit `gorm:"column:suit" json:"suit"`
		Rank      model.Rank `gorm:"column:rank" json:"rank"`
		Code      string     `gorm:"column:code" json:"code"`
		Remaining int        `gorm:"column:remaining" json:"remaining"`
	}
	views := make([]response, 0)
	err := c.db.Table("deck_cards").Select("deck_cards.deck_id, "+
		"decks.shuffled as shuffled, "+
		"decks.remaining as remaining, "+
		"cards.* as cards").
		Joins("INNER JOIN decks on deck_cards.deck_id = decks.id").
		Joins("INNER JOIN cards on deck_cards.card_id = cards.id").
		Where("deck_id = ? and deck_cards.deleted_at IS NULL", deckID).
		Limit(count).
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

func (c *cardRepositoryImpl) QueryInCards(ctx context.Context, v []string) ([]model.Card, database.Error) {
	cards := make([]model.Card, 0)
	err := c.db.Where("code IN ?", v).Find(&cards).Error
	if err != nil {
		log.Println("QueryInCards : err : ", err)
		return nil, database.NewError(err)
	}

	if len(cards) == 0 {
		return nil, database.NewError(gorm.ErrRecordNotFound)
	}

	return cards, nil
}

func (c *cardRepositoryImpl) QueryCards(ctx context.Context, params map[string]interface{}, filter string) ([]model.Card, database.Error) {
	cards := make([]model.Card, 0)
	err := c.db.Table("cards").
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

func NewCardRepositoryImpl(db *gorm.DB) repository.CardRepository {
	return &cardRepositoryImpl{db: db}
}
