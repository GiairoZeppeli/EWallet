package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type InitDbPostgres struct {
	db *sqlx.DB
}

func NewInitDbPostgres(db *sqlx.DB) *InitDbPostgres {
	return &InitDbPostgres{db: db}
}

func (r *InitDbPostgres) initDb() error {
	for i := 0; i < 10; i++ {
		var wallet_from, wallet
		query := fmt.Sprintf("INSERT INTO %s (wallet_from, wallet_to, amount) values ($1, $2, $3)", walletsTable)
		row := r.db.QueryRow(query,)
		if err := row.Scan(&balance); err != nil {
			return 0, nil
		}
	}
	return nil
}
