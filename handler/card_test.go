package handler

import (
	`net/http`
	`net/http/httptest`
	`strings`
	`testing`

	`github.com/gin-gonic/gin`
	`github.com/stretchr/testify/assert`

	`deck/core/api`
	`deck/core/config`
	`deck/mocks`
	`deck/model`
)

func TestCardHandler_CreateDecks(t *testing.T) {
	cases := map[string]struct {
		payload            string
		servicePayload     api.DeckRequest
		expectedStatusCode int
	}{
		"success": {
			payload: `
				{
					"partial": true,
					"cards": []{"AS"}
				}
			`,
			servicePayload: api.DeckRequest{
				Partial: true,
				Cards:   []string{"AS"},
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			body := strings.NewReader(tc.payload)
			c.Request, _ = http.NewRequest("POST", "/decks", body)

			cardService := new(mocks.Card)

			cardService.On(
				"CreateDeck",
				c,
				tc.servicePayload,
			).Return(
				model.Deck{
					Shuffled:  false,
					Remaining: 1,
				},
				nil,
			)

			cardHandler := NewCardHandler(cardService, &config.Config{})
			cardHandler.CreateDecks(c)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}

func TestCardHandler_DrawCards(t *testing.T) {
	cases := map[string]struct {
		payload            string
		servicePayload     api.DeckRequest
		expectedStatusCode int
	}{
		"success": {
			expectedStatusCode: http.StatusOK,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			body := strings.NewReader(tc.payload)
			c.Request, _ = http.NewRequest("PATCH", "/decks/deckId?count=1", body)

			cardService := new(mocks.Card)

			cardService.On(
				"GetDeckView",
				c,
				1,
				"deckId",
			).Return(
				model.DeckView{
					Shuffled:  false,
					Remaining: 1,
					Cards: []model.CardView{{
						Suit: model.Spade.String(),
						Rank: model.Ace.String(),
						Code: "AS",
					}},
				},
				nil,
			)

			cardService.On(
				"DrawCards",
				c,
				"deckId",
				[]string{"AS"},
			).Return(
				nil,
			)

			cardHandler := NewCardHandler(cardService, &config.Config{})
			cardHandler.DrawCards(c)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}

func TestCardHandler_GetDeck(t *testing.T) {
	cases := map[string]struct {
		payload            string
		servicePayload     api.DeckRequest
		expectedStatusCode int
	}{
		"success": {
			expectedStatusCode: http.StatusOK,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			body := strings.NewReader(tc.payload)
			c.Request, _ = http.NewRequest("PATCH", "/decks/deckId", body)

			cardService := new(mocks.Card)

			cardService.On(
				"GetDeckView",
				c,
				1,
				"deckId",
			).Return(
				model.DeckView{
					Shuffled:  false,
					Remaining: 1,
					Cards: []model.CardView{{
						Suit: model.Spade.String(),
						Rank: model.Ace.String(),
						Code: "AS",
					}},
				},
				nil,
			)

			cardHandler := NewCardHandler(cardService, &config.Config{})
			cardHandler.GetDeck(c)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}
