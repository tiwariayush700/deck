package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, out interface{}, uow *UnitOfWork) error
	Get(ctx context.Context, out interface{}, id interface{}, uow *UnitOfWork) error
	Update(ctx context.Context, out interface{}, id interface{}, uow *UnitOfWork) error
	Delete(ctx context.Context, out interface{}, id interface{}, uow *UnitOfWork) error //soft delete
}

// UnitOfWork represents a connection
type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

// NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
	if readOnly {
		return &UnitOfWork{DB: db, committed: false, readOnly: true} //normal db
	}

	return &UnitOfWork{DB: db.Begin(), committed: false, readOnly: false} //db with tx
}

// Complete marks end of unit of work
func (uow *UnitOfWork) Complete() {
	if !uow.committed && !uow.readOnly {
		uow.DB.Rollback()
	}
}

// Commit the transaction
func (uow *UnitOfWork) Commit() error {
	if !uow.readOnly {
		return uow.DB.Commit().Error
	}

	uow.committed = true

	return nil
}
