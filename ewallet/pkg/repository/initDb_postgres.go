package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type InitDbPostgres struct {
	db *sqlx.DB
}

func NewInitDbPostgres(db *sqlx.DB) *InitDbPostgres {
	return &InitDbPostgres{db: db}
}

func (r *InitDbPostgres) InitDb() error {
	var isEmpty bool
	query := fmt.Sprintf("SELECT (SELECT count(*) FROM %s) = 0", walletsTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&isEmpty); err != nil {
		return errors.New("cant check db to empty")
	}
	if !isEmpty {
		for i := 0; i < 9; i++ {
			s := strconv.FormatInt(int64(i), 10)
			h := sha1.New()
			h.Write([]byte(s))
			address := hex.EncodeToString(h.Sum(nil))
			query := fmt.Sprintf("INSERT INTO %s (address, amount) values ($1, $2)", walletsTable)
			r.db.Exec(query, address, 100.0)
		}
	}
	return nil
}
