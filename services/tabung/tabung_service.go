package tabung_service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	mdl "github.com/rbennum/service-account/models/tabung"
	repo "github.com/rbennum/service-account/repos/accounts"
	"github.com/rs/zerolog"
)

type TabungService struct {
	repo   repo.AccountRepo
	logger zerolog.Logger
}

func New(repo repo.AccountRepo, logger zerolog.Logger) TabungService {
	return TabungService{
		repo:   repo,
		logger: logger,
	}
}

func (ts TabungService) DepositBalance(ctx context.Context, depositedBalance int, accountNum string) (mdl.ResponseBody, error) {
	currentBalance, err := ts.repo.GetCurrentBalance(ctx, accountNum)
	if err != nil {
		return handlePGError(err)
	}

	currentBalance, err = ts.repo.UpdateAccountBalance(ctx, accountNum, currentBalance+depositedBalance)
	if err != nil {
		return handlePGError(err)
	}

	return mdl.ResponseBody{
		StatusCode: http.StatusOK,
		Balance:    currentBalance,
	}, nil
}

func handlePGError(err error) (mdl.ResponseBody, error) {
	if errors.Is(err, pgx.ErrNoRows) {
		return mdl.ResponseBody{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: pgx.ErrNoRows.Error(),
		}, pgx.ErrNoRows
	}
	return mdl.ResponseBody{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: err.Error(),
	}, err
}
