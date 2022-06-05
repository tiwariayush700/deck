package impl

import (
	`deck/repository`
)

type cardRepositoryImpl struct {
	repositoryImpl
}

func NewCardRepositoryImpl() repository.CardRepository {
	return &cardRepositoryImpl{}
}
