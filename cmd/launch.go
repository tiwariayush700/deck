package main

import (
	"context"
	`fmt`
	`log`

	"github.com/gin-gonic/gin"
	`gorm.io/driver/postgres`
	`gorm.io/gorm`

	"deck/core/config"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	configuration := config.GetAppConfiguration()

	//gin router
	router := gin.Default()

	db, err := gorm.Open(postgres.Open(postgresURI(configuration.PGConfig)),
		&gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalf("unable to connect to db : err %v", err)
	}

	app := newApp(router, configuration, db)

	err = app.autoMigrate()

	if err != nil {
		log.Fatalf("failed to autoMigrate with err %v", err)
	}

	app.start(ctx)
}

func postgresURI(pgConfig config.PGConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgConfig.PostgresServer, pgConfig.PostgresUser,
		pgConfig.PostgresPassword, pgConfig.PostgresDB, pgConfig.PostgresPort)
}
