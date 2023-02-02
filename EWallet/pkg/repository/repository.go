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
}

type Repository struct {
	Send
	GetLast
	GetBalance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		GetLast: NewSendPostgres(db),
	}
}
