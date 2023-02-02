package repository

import "github.com/jmoiron/sqlx"

type Send interface {
	Send()
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
	return &Repository{}
}
