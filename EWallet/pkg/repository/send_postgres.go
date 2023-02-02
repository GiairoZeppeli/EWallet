package repository

import (
	"errors"
	"fmt"
	"github.com/GiairoZeppeli/EWallet"
	"github.com/jmoiron/sqlx"
)

type SendPostgres struct {
	db *sqlx.DB
}

func NewSendPostgres(db *sqlx.DB) *SendPostgres {
	return &SendPostgres{db: db}
}

func (r *SendPostgres) Send(transaction EWallet.Transaction) error {
	var from_amount float32

	query := fmt.Sprintf("SELECT amount FROM %s WHERE address=$1 RETURNING amount", walletsTable)
	row := r.db.QueryRow(query, transaction.From)
	row.Scan(&from_amount)

	if err := from_amount < transaction.Amount; err != true {
		return errors.New("Insufficient Funds")
	}
	query = fmt.Sprintf("UPDATE %s SET amount = amount+$1 WHERE address=$2", walletsTable)
	row = r.db.QueryRow(query, transaction.From)
	r.db.Exec(query, transaction.Amount, transaction.To)

	query = fmt.Sprintf("INSERT INTO %s (wallet_from, wallet_to, amount) values ($1, $2, $3)", transactionsTable)
	r.db.Exec(query, transaction.From, transaction.To, transaction.Amount)
	return nil
}
