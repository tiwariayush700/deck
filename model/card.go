//go:generate stringer -type=Suit,Rank

package model

import (
	`time`

	`github.com/hashicorp/go-uuid`

	`deck/core/database`
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

var Suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	MinRank = Ace
	MaxRank = King
)

type Card struct {
	database.BaseModel
	Suit Suit   `json:"suit" gorm:"suit:text;not null"`
	Rank Rank   `json:"value"`
	Code string `json:"code"`
}

type Deck struct {
	database.BaseModel
	Shuffled  bool `json:"shuffled"`
	Remaining int  `json:"remaining"`
}

func NewDeck(shuffled bool, remaining int) *Deck {
	//deckID, _ := uuid.GenerateUUID()

	return &Deck{
		//BaseModel: database.BaseModel{
		//	ID:        deckID,
		//	CreatedAt: time.Now(),
		//	UpdatedAt: time.Now(),
		//},
		Shuffled:  shuffled,
		Remaining: remaining,
	}
}

type DeckCard struct {
	database.BaseModel
	SequenceID int `gorm:"autoIncrement:true;type:bigserial"`

	//foreign keys
	DeckID string `json:"deck_id"`
	CardID string `json:"card_id"`

	Deck *Deck `json:"-"`
	Card *Card `json:"-"`
}

func NewDeckCard(deckID string, cardID string) *DeckCard {
	deckCardID, _ := uuid.GenerateUUID()

	return &DeckCard{
		BaseModel: database.BaseModel{
			ID:        deckCardID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DeckID: deckID,
		CardID: cardID,
	}
}

type CardView struct {
	Suit string `json:"suit"`
	Rank string `json:"value"`
	Code string `json:"code"`
}

type DeckView struct {
	DeckID    string     `gorm:"column:deck_id" json:"deck_id"`
	Shuffled  bool       `gorm:"column:shuffled" json:"shuffled"`
	Remaining int        `gorm:"column:remaining" json:"remaining"`
	Cards     []CardView `gorm:"column:cards" json:"cards"`
}
