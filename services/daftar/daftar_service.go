package daftar_service

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	mdl "github.com/rbennum/service-account/models/daftar"
	"github.com/rbennum/service-account/models/entity"
	rp "github.com/rbennum/service-account/repos/daftar"
	"github.com/rs/zerolog"
)

type DaftarService struct {
	repo   rp.DaftarRepo
	logger zerolog.Logger
}

func New(repo rp.DaftarRepo, logger zerolog.Logger) DaftarService {
	return DaftarService{
		repo:   repo,
		logger: logger,
	}
}

func (d DaftarService) CreateCustomer(ctx context.Context, input mdl.RequestBody) (mdl.ResponseBody, error) {
	entity := entity.CustomerEntity{
		Name:     input.Name,
		NIK:      input.ID,
		PhoneNum: input.Phone,
	}
	err := d.repo.CreateCustomer(ctx, entity)
	if err != nil {
		d.logger.Debug().Err(err).Msg("debugging error")
		return handlePGError(err)
	}

	accountNum, err := d.repo.GenerateAccountNumber(ctx)
	if err != nil {
		return mdl.ResponseBody{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}, err
	}

	err = d.repo.CreateAccount(ctx, input.ID, accountNum)
	if err != nil {
		return handlePGError(err)
	}

	return mdl.ResponseBody{
		StatusCode: http.StatusOK,
		Account:    accountNum,
	}, nil
}

func handlePGError(err error) (mdl.ResponseBody, error) {
	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
		return mdl.ResponseBody{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Duplicate entry: " + pgErr.Detail,
		}, pgErr
	}

	return mdl.ResponseBody{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: err.Error(),
	}, err
}
