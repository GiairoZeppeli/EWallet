package repository

import (
	"fmt"
	"github.com/GiairoZeppeli/EWallet"
	"github.com/jmoiron/sqlx"
)

type GetLastPostgres struct {
	db *sqlx.DB
}

func NewGetLastPostgres(db *sqlx.DB) *GetLastPostgres {
	return &GetLastPostgres{db: db}
}

func (r *GetLastPostgres) GetLast(count int) ([]EWallet.Transaction, error) {
	var transactions []EWallet.Transaction

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT $1", transactionsTable)
	err := r.db.Select(&transactions, query, count)

	return transactions, err
}
