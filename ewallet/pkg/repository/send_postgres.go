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

func (r *SendPostgres) Send(transaction ewallet.Transaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var from_amount float32
	getFromAmountQuery := fmt.Sprintf("SELECT amount FROM %s WHERE address=$1", walletsTable)
	row := tx.QueryRow(getFromAmountQuery, transaction.From)
	if err := row.Scan(&from_amount); err != nil {
		tx.Rollback()
		return errors.New("cant read balanceFrom")
	}

	if from_amount < transaction.Amount {
		return errors.New("Insufficient Funds")
	}
	updateToBalanceQuery := fmt.Sprintf("UPDATE %s SET amount = amount+$1 WHERE address=$2", walletsTable)
	_, err = tx.Exec(updateToBalanceQuery, transaction.Amount, transaction.To)
	if err != nil {
		tx.Rollback()
		return errors.New("cant update balance")
	}
	updateFromBalanceQuery := fmt.Sprintf("UPDATE %s SET amount = amount-$1 WHERE address=$2", walletsTable)
	_, err = tx.Exec(updateFromBalanceQuery, transaction.Amount, transaction.From)
	if err != nil {
		tx.Rollback()
		return errors.New("cant update balance")
	}

	createTransactionQuery := fmt.Sprintf("INSERT INTO %s (wallet_from, wallet_to, amount) values ($1, $2, $3)", transactionsTable)
	_, err = tx.Exec(createTransactionQuery, transaction.From, transaction.To, transaction.Amount)
	if err != nil {
		tx.Rollback()
		return errors.New("cant add transaction into database")
	}
	tx.Commit()
	return nil
}
