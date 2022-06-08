package main

import (
	"context"
	`log`

	"github.com/gin-gonic/gin"

	"deck/core/config"
	`deck/repository/impl`
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	configuration := config.GetAppConfiguration()

	//gin router
	router := gin.Default()

	repository, pgDB, err := impl.NewRepositoryImpl(configuration.PGConfig)
	if err != nil {
		log.Fatalf("unable to connect to db : err %v", err)
	}

	app := newApp(router, configuration, pgDB, repository)

	err = app.autoMigrate()

	if err != nil {
		log.Fatalf("failed to autoMigrate with err %v", err)
	}

	app.start(ctx)
}
