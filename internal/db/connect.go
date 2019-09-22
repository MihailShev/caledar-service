package db

import "github.com/jmoiron/sqlx"

type Config struct {
	Dns string
}

func connect(dns string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", dns)

	if err != nil {
		return db, err
	}

	err = db.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
}
