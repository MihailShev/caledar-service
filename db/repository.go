package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

func Connect() {
	dns := "postgres://mshev:123qwe@localhost:5432/calendar?sslmode=disable"
	db, err := sqlx.Open("pgx", dns)

	if err != nil {
		log.Println(err)
	}

	err = db.Ping()

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("connect ot db calendar is established")
	}
}
