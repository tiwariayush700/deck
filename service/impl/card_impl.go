package impl

import (
	`context`
	`log`

	`deck/core/api`
	`deck/core/database`
	`deck/model`
	`deck/repository`
	`deck/service`
	`deck/util`
)

type cardServiceImpl struct {
	CardRepository repository.CardRepository
}

func (c *cardServiceImpl) DrawCards(ctx context.Context, deckID string, cardIDs []string) error {

	dbErr := c.CardRepository.DeleteCardsFromDeck(ctx, deckID, cardIDs)

	if dbErr != nil {
		if dbErr.IsRecordNotFoundError() {
			return api.NewHTTPError(api.ErrorCodeResourceNotFound, "No deck found for the provided id")
		}
		return api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Failed to draw cards")
	}

	return nil
}

func (c *cardServiceImpl) GetDeckView(ctx context.Context, count int, deckID string) (*model.DeckView, error) {

	deckView, dbErr := c.CardRepository.GetDeckView(ctx, deckID, count)

	if dbErr != nil {
		if dbErr.IsRecordNotFoundError() {
			return nil, api.NewHTTPError(api.ErrorCodeResourceNotFound, "No deck found for the provided id")
		}
		return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Failed to query decks")
	}

	if len(deckView.Cards) == 0 {
		return nil, api.NewHTTPError(api.ErrorCodeResourceNotFound, "Sorry, all cards are already drawn out of this deck")
	}

	return deckView, nil
}

func (c *cardServiceImpl) CreateDeck(ctx context.Context, request api.DeckRequest) (*model.Deck, error) {

	//User can not have both partial and shuffle
	if request.Partial && request.Shuffle {
		return nil, api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "You cannot have partial and shuffle both")
	}

	cards := make([]model.Card, 0)
	params := make(map[string]interface{})
	var dbErr database.Error
	if request.Partial {
		if len(request.Cards) == 0 {
			return nil, api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "Please provide cards if you want a partial deck")
		}
		cards, dbErr = c.CardRepository.QueryInCards(ctx, request.Cards)
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
		cards, dbErr = c.CardRepository.QueryCards(ctx, params, "")

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
	dbErr = c.CardRepository.CreateDeckCards(ctx, deck, cards)

	if dbErr != nil {
		log.Printf("Failed to Create deck : err %v", dbErr)
		return nil, api.NewHTTPError(api.ErrorCodeDatabaseFailure, "Something went wrong while creating deck")
	}

	return deck, nil
}

func NewCardServiceImpl(cardRepository repository.CardRepository) service.Card {
	return &cardServiceImpl{CardRepository: cardRepository}
}
