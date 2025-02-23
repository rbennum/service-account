package check_service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	tm "github.com/rbennum/service-account/models/tabung"
	repo "github.com/rbennum/service-account/repos/accounts"
	"github.com/rs/zerolog"
)

type CheckService struct {
	repo   repo.AccountRepo
	logger zerolog.Logger
}

func New(repo repo.AccountRepo, logger zerolog.Logger) CheckService {
	return CheckService{
		repo:   repo,
		logger: logger,
	}
}

func (cs CheckService) CheckBalance(ctx context.Context, accountNum string) (tm.ResponseBody, error) {
	currentBalance, err := cs.repo.GetCurrentBalance(ctx, accountNum)
	if err != nil {
		return handlePGError(err)
	}

	return tm.ResponseBody{
		StatusCode: http.StatusOK,
		Balance:    currentBalance,
	}, nil
}

func handlePGError(err error) (tm.ResponseBody, error) {
	if errors.Is(err, pgx.ErrNoRows) {
		return tm.ResponseBody{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: pgx.ErrNoRows.Error(),
		}, pgx.ErrNoRows
	}
	return tm.ResponseBody{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: err.Error(),
	}, err
}
