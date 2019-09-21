package repository

import (
	"context"
	"github.com/MihailShev/calendar-service/internal/calendar"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}

type Repository struct {
	db     *sqlx.DB
	logger Logger
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

func NewRepository(logger Logger) (*Repository, error) {
	db, err := connect()

	if err != nil {
		return &Repository{}, err
	}

	logger.Infof("Connect to calendar db is established\n")

	return &Repository{db: db, logger: logger}, err
}

func (r *Repository) CreateEvent(ctx context.Context, e calendar.Event) (int64, error) {
	var uuid int64

	query := `INSERT INTO event(user_id, title, description, start, "end", notice_time)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid`

	err := r.db.QueryRowContext(ctx, query, e.UserId, e.Title, e.Description, e.Start, e.End, e.NotifyTime).Scan(&uuid)

	return uuid, err
}

func (r *Repository) GetEventById(ctx context.Context, uuid int64) (calendar.Event, error) {
	var event calendar.Event
	query := `SELECT * FROM event WHERE uuid = :uuid;`
	rows, err := r.db.NamedQueryContext(ctx, query, map[string]interface{}{"uuid": uuid})

	defer r.closeRows(rows)

	if err != nil {
		return event, err
	}

	rows.Next()
	err = rows.Scan(
		&event.UUID,
		&event.UserId,
		&event.Title,
		&event.Description,
		&event.Start,
		&event.End,
		&event.NotifyTime)

	return event, err
}

func (r *Repository) UpdateEvent(ctx context.Context, event calendar.Event) (calendar.Event, error) {
	var updated calendar.Event

	query := `UPDATE event 
		SET (user_id, title, description, start, "end", notice_time) = 
			(:userId, :title, :description, :start, :end, :noticeTime)
		WHERE uuid = :uuid
		RETURNING uuid;`

	rows, err := r.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"uuid":        event.UUID,
		"userId":      event.UserId,
		"title":       event.Title,
		"description": event.Description,
		"start":       event.Start,
		"end":         event.End,
		"noticeTime":  event.NotifyTime,
	})

	defer r.closeRows(rows)

	if err != nil {
		return updated, err
	}
	var uuid int64

	rows.Next()
	err = rows.Scan(&uuid)

	if err != nil {
		return updated, err
	}

	updated, err = r.GetEventById(ctx, uuid)

	return updated, err
}

func (r *Repository) closeRows(rows *sqlx.Rows) {
	err := rows.Close()

	if err != nil {
		r.logger.Errorf("%s", err.Error())
	}
}
