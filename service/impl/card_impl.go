package impl

import (
	`context`
	`log`

	`gorm.io/gorm`

	`deck/core/api`
	`deck/core/database`
	`deck/model`
	`deck/repository`
	`deck/service`
	`deck/util`
)

type cardServiceImpl struct {
	CardRepository repository.CardRepository
	DB             *gorm.DB
}

func (c *cardServiceImpl) GetDeckView(ctx context.Context, count int, deckID string) (*model.DeckView, error) {

	uow := repository.NewUnitOfWork(c.DB, true)
	defer uow.Complete()

	deckView, dbErr := c.CardRepository.GetDeckView(ctx, deckID, count, uow)

	if dbErr != nil {
		if dbErr.IsRecordNotFoundError() {
			return nil, api.NewHTTPError(api.ErrorCodeResourceNotFound, "No deck found for the provided id")
		}
		return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Failed to query decks")
	}

	err := uow.Commit()

	if err != nil {
		return nil, api.NewHTTPError(api.ErrorCodeInternalError, "Something went wrong")
	}

	return deckView, nil
}

func (c *cardServiceImpl) CreateDeck(ctx context.Context, request api.DeckRequest) (*model.Deck, error) {

	//User can not have both partial and shuffle
	if request.Partial && request.Shuffle {
		return nil, api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "You cannot have partial and shuffle both")
	}

	uow := repository.NewUnitOfWork(c.DB, false)
	defer uow.Complete()

	cards := make([]model.Card, 0)
	params := make(map[string]interface{})
	var dbErr database.Error
	if request.Partial {
		if len(request.Cards) == 0 {
			return nil, api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "Please provide cards if you want a partial deck˘")
		}
		cards, dbErr = c.CardRepository.QueryInCards(ctx, request.Cards, uow)
		if dbErr != nil {
			if dbErr.IsRecordNotFoundError() {
				return nil, api.NewHTTPError(api.ErrorCodeResourceNotFound, "No cards found for the provided codes")
			}
			log.Printf("Failed to QueryInCards : err %v", dbErr)
			return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Something went wrong")
		}

		if len(cards) != len(request.Cards) {
			return nil, api.NewHTTPError(api.ErrorCodeResourceNotFound, "No cards found for the provided codes")
		}
	} else {
		cards, dbErr = c.CardRepository.QueryCards(ctx, params, "", uow)
		if dbErr != nil {
			if dbErr.IsRecordNotFoundError() {
				return nil, api.NewHTTPError(api.ErrorCodeResourceNotFound, "No cards found for the provided codes")
			}
			log.Printf("Failed to QueryCards : err %v", dbErr)
			return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Something went wrong")
		}
	}

	if request.Shuffle {
		cards = util.Shuffle(cards)
	}

	deck := model.NewDeck(request.Shuffle, len(cards))
	dbErr = c.CardRepository.Create(ctx, deck, uow)
	if dbErr != nil {
		log.Printf("Failed to Create deck : err %v", dbErr)
		return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Something went wrong while creating deck")
	}

	deckCards := make([]model.DeckCard, 0)
	for _, card := range cards {
		deckCards = append(deckCards, *model.NewDeckCard(deck.ID, card.ID))
	}

	dbErr = c.CardRepository.Create(ctx, &deckCards, uow)
	if dbErr != nil {
		log.Printf("Failed to Create deckCards : err %v", dbErr)
		return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Something went wrong while creating deck cards")
	}

	err := uow.Commit()
	if err != nil {
		log.Printf("Failed to commit unit of work : err %v", dbErr)
		return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Something went wrong")
	}

	return deck, nil
}

func NewCardServiceImpl(cardRepository repository.CardRepository, db *gorm.DB) service.Card {
	return &cardServiceImpl{CardRepository: cardRepository, DB: db}
}
