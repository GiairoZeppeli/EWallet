package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GetBalancePostgres struct {
	db *sqlx.DB
}

func NewGetBalancePostgres(db *sqlx.DB) *GetBalancePostgres {
	return &GetBalancePostgres{db: db}
}

func (r *GetBalancePostgres) GetBalance(address string) (float32, error) {
	var balance float32
	query := fmt.Sprintf("SELECT amount FROM %s WHERE address=%s RETURNING balance", walletsTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&balance); err != nil {
		return 0, errors.New("cant get balance from database")
	}
	return balance, nil
}
