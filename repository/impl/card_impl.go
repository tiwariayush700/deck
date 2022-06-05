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
