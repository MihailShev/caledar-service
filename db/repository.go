package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func connect() (*sqlx.DB, error) {
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

func NewRepository() (Repository, error) {
	db, err := connect()

	if err != nil {
		return Repository{}, err
	}

	return Repository{db: db}, nil
}

func (r *Repository) CreateEvent(e EventModel) (int64, error) {
	var id int64
	query := `INSERT INTO event(user_id, title, description, start, "end", notice_time)
			VALUES ($1, $2, $3, $4, $5, $6)`

	res, err := r.db.Exec(query, e.UserId, e.Title, e.Description, e.Start, e.End, e.NoticeTime)

	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetEventById(id int64) (EventModel, error) {
	query := `SELECT * FROM WHERE uuid = $1;`
	var e EventModel

	row := r.db.QueryRow(query, id)
	err := row.Scan(&e)

	return e, err
}
