package tarik_service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	tm "github.com/rbennum/service-account/models/tabung"
	repo "github.com/rbennum/service-account/repos/accounts"
	"github.com/rs/zerolog"
)

type TarikService struct {
	repo   repo.AccountRepo
	logger zerolog.Logger
}

func New(repo repo.AccountRepo, logger zerolog.Logger) TarikService {
	return TarikService{
		repo:   repo,
		logger: logger,
	}
}

func (ts TarikService) WithdrawBalance(ctx context.Context, withdrawedBalance int, accountNum string) (tm.ResponseBody, error) {
	currentBalance, err := ts.repo.GetCurrentBalance(ctx, accountNum)
	if err != nil {
		return handlePGError(err)
	}

	totalBalance := currentBalance - withdrawedBalance
	if totalBalance < 0 {
		return tm.ResponseBody{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "total balance should never be negative",
		}, fmt.Errorf("total balance should never be negative")
	}

	currentBalance, err = ts.repo.UpdateAccountBalance(ctx, accountNum, totalBalance)
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
