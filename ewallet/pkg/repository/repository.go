package repository

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/jmoiron/sqlx"
)

type Send interface {
	Send(transaction ewallet.Transaction) error
}

type GetLast interface {
	GetLast(count int) ([]ewallet.Transaction, error)
}

type GetBalance interface {
	GetBalance(address string) (float32, error)
}
type InitDb interface {
	InitDb() error
}

type Repository struct {
	Send
	GetLast
	GetBalance
	InitDb
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Send:       NewSendPostgres(db),
		GetBalance: NewGetBalancePostgres(db),
		InitDb:     NewInitDbPostgres(db),
		GetLast:    NewGetLastPostgres(db),
	}
}
