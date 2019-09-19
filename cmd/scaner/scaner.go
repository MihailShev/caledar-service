package main

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Event struct {
	UUID        int64
	Title       string
	Start       time.Time
	End         time.Time
	NotifyTime  time.Time `db:"notice_time"`
	Description string
	UserId      uint64 `db:"user_id"`
}

func main() {

}

func dbConnect() (*sqlx.DB, error) {
	dns := "postgres://mshev:123qwe@localhost:5432/calendar?sslmode=disable"
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
