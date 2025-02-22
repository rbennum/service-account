package main

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rbennum/service-account/database/migrate"
	"github.com/rbennum/service-account/database/postgres"
	daftar_handler "github.com/rbennum/service-account/handlers/daftar"
	"github.com/rbennum/service-account/middleware"
	daftar_repo "github.com/rbennum/service-account/repos/daftar"
	daftar_service "github.com/rbennum/service-account/services/daftar"
	"github.com/rbennum/service-account/utils/config"
	log "github.com/rbennum/service-account/utils/log"
	"github.com/rs/zerolog"
)

func main() {
	err := log.Init()
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("error initiating logger")
	}

	config := config.GetConfig()

	db, err := postgres.New()
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("error initiating database connection")
	}
	defer db.Close()

	log.Logger.Info().Msg("logger: configured")
	err = migrate.Migrate()
	if err != nil {
		log.Logger.Fatal().Err(err).
			Msg("error migrating database")
	}

	e := echo.New()
	e.Use(middleware.RequestIDMiddleware())
	e.Use(middleware.LoggerMiddleware())
	setDaftar(e, db, log.Logger)
	log.Logger.Info().
		Msg(fmt.Sprintf("listening to port %s", config.Port))
	e.Start(":" + config.Port)
}

func setDaftar(e *echo.Echo, db *pgxpool.Pool, logger zerolog.Logger) {
	repo := daftar_repo.New(db)
	svc := daftar_service.New(repo, logger)
	handler := daftar_handler.New(svc, logger)
	e.POST("/daftar", handler.PostDaftar)
}
