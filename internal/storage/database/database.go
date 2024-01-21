package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	DB *sql.DB
	//TODO: add custom options
}

func NewPostgres(uri string) (*Postgres, error) {
	pg := &Postgres{}

	//// Custom options
	//for _, opt := range opts {
	//	opt(pg)
	//}

	db, err := sql.Open("pgx", uri)
	if err != nil {
		return nil, err
	}

	pg.DB = db
	return pg, nil
}
