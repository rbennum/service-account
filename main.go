package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rbennum/service-account/database/migrate"
	"github.com/rbennum/service-account/middleware"
	"github.com/rbennum/service-account/utils/config"
	log "github.com/rbennum/service-account/utils/log"
)

func main() {
	err := log.Init()
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("error initiating logger")
	}

	config := config.GetConfig()

	log.Logger.Info().Msg("logger: configured")
	err = migrate.Migrate()
	if err != nil {
		log.Logger.Fatal().Err(err).
			Msg("error migrating database")
	}

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.ResponseLogger())
	e.Start(":" + config.Port)
	log.Logger.Info().
		Msg(fmt.Sprintf("listening to port %s", config.Port))
}
