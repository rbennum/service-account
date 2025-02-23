package main

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rbennum/service-account/database/migrate"
	"github.com/rbennum/service-account/database/postgres"
	check_handler "github.com/rbennum/service-account/handlers"
	daftar_handler "github.com/rbennum/service-account/handlers/daftar"
	tabung_handler "github.com/rbennum/service-account/handlers/tabung"
	tarik_handler "github.com/rbennum/service-account/handlers/tarik"
	"github.com/rbennum/service-account/middleware"
	account_repo "github.com/rbennum/service-account/repos/accounts"
	user_repo "github.com/rbennum/service-account/repos/users"
	check_service "github.com/rbennum/service-account/services/check"
	daftar_service "github.com/rbennum/service-account/services/daftar"
	tabung_service "github.com/rbennum/service-account/services/tabung"
	tarik_service "github.com/rbennum/service-account/services/tarik"
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
	setTabung(e, db, log.Logger)
	setTarik(e, db, log.Logger)
	setCheckSaldo(e, db, log.Logger)
	log.Logger.Info().
		Msg(fmt.Sprintf("listening to port %s", config.Port))
	e.Start(":" + config.Port)
}

func setDaftar(e *echo.Echo, db *pgxpool.Pool, logger zerolog.Logger) {
	repo := user_repo.New(db)
	svc := daftar_service.New(repo, logger)
	handler := daftar_handler.New(svc, logger)
	e.POST("/daftar", handler.PostDaftar)
}

func setTabung(e *echo.Echo, db *pgxpool.Pool, logger zerolog.Logger) {
	repo := account_repo.New(db)
	svc := tabung_service.New(repo, logger)
	handler := tabung_handler.New(svc, logger)
	e.POST("/tabung", handler.DepositBalance)
}

func setTarik(e *echo.Echo, db *pgxpool.Pool, logger zerolog.Logger) {
	repo := account_repo.New(db)
	svc := tarik_service.New(repo, logger)
	handler := tarik_handler.New(svc, logger)
	e.POST("/tarik", handler.WithdrawBalance)
}

func setCheckSaldo(e *echo.Echo, db *pgxpool.Pool, logger zerolog.Logger) {
	repo := account_repo.New(db)
	svc := check_service.New(repo, logger)
	handler := check_handler.New(svc, logger)
	e.GET("/saldo/:no_rekening", handler.CheckBalance)
}
