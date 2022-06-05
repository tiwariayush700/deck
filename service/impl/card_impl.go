package impl

import (
	`gorm.io/gorm`

	`deck/repository`
	`deck/service`
)

type cardServiceImpl struct {
	CardRepository repository.CardRepository
	DB             *gorm.DB
}

func NewCardServiceImpl(cardRepository repository.CardRepository, db *gorm.DB) service.Card {
	return &cardServiceImpl{CardRepository: cardRepository, DB: db}
}
