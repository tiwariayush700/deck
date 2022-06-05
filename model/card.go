//go:generate stringer -type=Suit,Rank

package model

import (
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
	Suit Suit   `json:"suit" gorm:"type:text;check:type = 'SPADE' or type = 'DIAMOND' or type = 'CLUB' or type = 'HEART';not null"`
	Rank Rank   `json:"value"`
	Code string `json:"code"`
}

type Deck struct {
	database.BaseModel
	Shuffle bool `json:"shuffle"`
}

type DeckCard struct {
	database.BaseModel
	SequenceID int `gorm:"autoIncrement:true;type:bigserial"`

	//foreign keys
	DeckID string `json:"deck_id"`
	CardID string `json:"card_id"`

	Topup *Deck `json:"-"`
	Card  *Card `json:"-"`
}
