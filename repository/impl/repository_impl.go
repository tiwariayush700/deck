package impl

import (
	`context`

	`deck/core/database`
	"deck/repository"
)

type repositoryImpl struct {
}

func NewRepository() repository.Repository {
	return &repositoryImpl{}
}

func (impl *repositoryImpl) Create(ctx context.Context, out interface{}, uow *repository.UnitOfWork) database.Error {
	err := uow.DB.Create(out).Error
	return database.NewError(err)
}

func (impl *repositoryImpl) Get(ctx context.Context, out interface{}, id interface{}, uow *repository.UnitOfWork) database.Error {
	err := uow.DB.First(out, "id = ?", id).Error
	return database.NewError(err)

}

func (impl *repositoryImpl) Update(ctx context.Context, out interface{}, id interface{}, uow *repository.UnitOfWork) database.Error {

	err := uow.DB.Where("id = ?", id).
		Updates(out).Error
	if err != nil {
		return database.NewError(err)
	}

	return nil

}

// Delete soft delete
func (impl *repositoryImpl) Delete(ctx context.Context, out interface{}, id interface{}, uow *repository.UnitOfWork) database.Error {

	err := uow.DB.Where(" id = ? ", id).Delete(out).Error
	if err != nil {
		return database.NewError(err)
	}

	return nil
}
