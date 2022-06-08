package impl

import (
	`context`
	`fmt`
	`log`

	`gorm.io/driver/postgres`
	`gorm.io/gorm`

	`deck/core/config`
	`deck/core/database`
	"deck/repository"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(pgConfig config.PGConfig) (repository.Repository, *gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(postgresURI(pgConfig)),
		&gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Printf("Failed to init new repo impl with err : %v", err)

		return nil, nil, err
	}

	return &repositoryImpl{db: db}, db, nil
}

func postgresURI(pgConfig config.PGConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgConfig.PostgresServer, pgConfig.PostgresUser,
		pgConfig.PostgresPassword, pgConfig.PostgresDB, pgConfig.PostgresPort)
}

func (impl *repositoryImpl) Create(ctx context.Context, out interface{}) database.Error {
	err := impl.db.Create(out).Error

	return database.NewError(err)
}

func (impl *repositoryImpl) Get(ctx context.Context, out interface{}, id interface{}) database.Error {
	err := impl.db.First(out, "id = ?", id).Error

	return database.NewError(err)
}

func (impl *repositoryImpl) Update(ctx context.Context, out interface{}, id interface{}) database.Error {

	err := impl.db.Where("id = ?", id).
		Updates(out).Error

	if err != nil {
		return database.NewError(err)
	}

	return nil

}

// Delete soft delete
func (impl *repositoryImpl) Delete(ctx context.Context, out interface{}, id interface{}) database.Error {

	err := impl.db.Where(" id = ? ", id).Delete(out).Error

	if err != nil {
		return database.NewError(err)
	}

	return nil
}
