package util

import (
	`fmt`
	`testing`
	`time`

	`github.com/stretchr/testify/assert`

	`deck/core/database`
	`deck/model`
)

func TestGetCode(t *testing.T) {

	code := GetCode(model.Spade, model.King)
	expectedCode := "KS"

	assert.Equal(t, expectedCode, code)
}

func TestShuffle(t *testing.T) {

	cards := Shuffle([]model.Card{})
	for _, suit := range model.Suits {
		for rank := model.MinRank; rank <= model.MaxRank; rank++ {
			cardID := GetCode(suit, rank)
			code := GetCode(suit, rank)
			card := model.Card{
				BaseModel: database.BaseModel{
					ID:        cardID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Suit: suit,
				Rank: rank,
				Code: code,
			}
			cards = append(cards, card)
		}
	}

	shuffledCards := Shuffle(cards)

	fmt.Println(shuffledCards)

	assert.NotEqual(t, shuffledCards[0].Code, "AS")
}
