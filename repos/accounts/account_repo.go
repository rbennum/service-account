package account_repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepo struct {
	db *pgxpool.Pool
}

func New(database *pgxpool.Pool) AccountRepo {
	return AccountRepo{
		db: database,
	}
}

func (ar AccountRepo) GetCurrentBalance(ctx context.Context, accountNum string) (int, error) {
	query := `
		SELECT balance 
		FROM accounts 
		WHERE account_num = $1
	`
	var balance int
	err := ar.db.QueryRow(ctx, query, accountNum).
		Scan(&balance)
	if err != nil {
		return -1, err
	}
	return balance, nil
}

func (ar AccountRepo) UpdateAccountBalance(ctx context.Context, accountNum string, newBalance int) (int, error) {
	tx, err := ar.db.Begin(ctx)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE accounts 
		SET balance = $1 
		WHERE account_num = $2 
		RETURNING balance
	`
	var balance int
	err = tx.QueryRow(ctx, query, newBalance, accountNum).
		Scan(&balance)
	if err != nil {
		return -1, err
	}
	return balance, tx.Commit(ctx)
}
