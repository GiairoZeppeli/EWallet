package repository

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/jmoiron/sqlx"
)

type Send interface {
	Send(transaction EWallet.Transaction) error
}

type GetLast interface {
}

type GetBalance interface {
	GetBalance(address string) (float32, error)
}

type Repository struct {
	Send
	GetLast
	GetBalance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Send:       NewSendPostgres(db),
		GetBalance: NewGetBalancePostgres(db),
	}
}
