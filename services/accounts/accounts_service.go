package accounts_service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	tm "github.com/rbennum/service-account/models/accounts"
	repo "github.com/rbennum/service-account/repos/accounts"
	"github.com/rs/zerolog"
)

type AccountsService struct {
	repo   repo.AccountRepo
	logger zerolog.Logger
}

func New(repo repo.AccountRepo, logger zerolog.Logger) AccountsService {
	return AccountsService{
		repo:   repo,
		logger: logger,
	}
}

func (as AccountsService) CheckBalance(ctx context.Context, accountNum string) (tm.ResponseBody, error) {
	currentBalance, err := as.repo.GetCurrentBalance(ctx, accountNum)
	if err != nil {
		return handlePGError(err)
	}

	return tm.ResponseBody{
		StatusCode: http.StatusOK,
		Balance:    &currentBalance,
	}, nil
}

func (as AccountsService) DepositBalance(ctx context.Context, depositedBalance int, accountNum string) (tm.ResponseBody, error) {
	currentBalance, err := as.repo.GetCurrentBalance(ctx, accountNum)
	if err != nil {
		return handlePGError(err)
	}

	currentBalance, err = as.repo.UpdateAccountBalance(ctx, accountNum, currentBalance+depositedBalance)
	if err != nil {
		return handlePGError(err)
	}

	return tm.ResponseBody{
		StatusCode: http.StatusOK,
		Balance:    &currentBalance,
	}, nil
}

func (as AccountsService) WithdrawBalance(ctx context.Context, withdrawedBalance int, accountNum string) (tm.ResponseBody, error) {
	currentBalance, err := as.repo.GetCurrentBalance(ctx, accountNum)
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

	currentBalance, err = as.repo.UpdateAccountBalance(ctx, accountNum, totalBalance)
	if err != nil {
		return handlePGError(err)
	}

	return tm.ResponseBody{
		StatusCode: http.StatusOK,
		Balance:    &currentBalance,
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
