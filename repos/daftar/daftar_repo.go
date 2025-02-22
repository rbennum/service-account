package daftar_repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	e "github.com/rbennum/service-account/models/entity"
)

type DaftarRepo struct {
	db *pgxpool.Pool
}

func New(database *pgxpool.Pool) DaftarRepo {
	return DaftarRepo{
		db: database,
	}
}

func (d DaftarRepo) CreateCustomer(ctx context.Context, entity e.CustomerEntity) error {
	query := `
		INSERT INTO users(nik, name, phone_num)
		VALUES($1, $2, $3)
	`
	_, err := d.db.Exec(
		ctx,
		query,
		entity.NIK, entity.Name, entity.PhoneNum,
	)
	fmt.Printf("%+v\n", err)
	return err
}

func (d DaftarRepo) CreateAccount(ctx context.Context, nik string, accountNum string) error {
	query := `
		INSERT INTO accounts(account_num, nik, balance)
		VALUES($1, $2, $3)
		RETURNING account_num;
	`
	_, err := d.db.Exec(ctx, query, accountNum, nik, 0)
	return err
}

func (d DaftarRepo) GenerateAccountNumber(ctx context.Context) (string, error) {
	var accountNumber int64
	query := `
		SELECT nextval('account_number_seq')
	`
	err := d.db.QueryRow(ctx, query).Scan(&accountNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%010d", accountNumber), nil
}
