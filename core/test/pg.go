package test

import (
	`fmt`
	`log`

	`gorm.io/driver/postgres`
	`gorm.io/gorm`
)

func DB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(postgresURI()),
		&gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Printf("Failed to init new repo impl with err : %v", err)

		return nil, err
	}

	return db, nil
}

func postgresURI() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "deck_user",
		"Vm28xMykxKM", "deck_db", "5431")
}
