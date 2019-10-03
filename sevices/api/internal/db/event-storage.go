package db

import (
	"context"
	"github.com/MihailShev/calendar-service/pkg/connector"
	"github.com/MihailShev/calendar-service/sevices/api/internal/calendar"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}

type EventStorage struct {
	db     *sqlx.DB
	logger Logger
}

func NewEventStorage(logger Logger, config connector.Config) (*EventStorage, error) {
	db, err := connector.Connect(config.Dns)

	if err != nil {
		return &EventStorage{}, err
	}

	logger.Infof("Connect to calendar db is established\n")

	return &EventStorage{db: db, logger: logger}, err
}

func (s *EventStorage) CreateEvent(ctx context.Context, e calendar.Event) (int64, error) {
	var uuid int64

	query := `INSERT INTO event(user_id, title, description, start, "end", notice_time)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid`

	err := s.db.QueryRowContext(ctx, query, e.UserId, e.Title, e.Description, e.Start, e.End, e.NotifyTime).Scan(&uuid)

	return uuid, err
}

func (s *EventStorage) GetEventById(ctx context.Context, uuid int64) (calendar.Event, error) {
	var event calendar.Event
	query := `SELECT * FROM event WHERE uuid = :uuid;`
	rows, err := s.db.NamedQueryContext(ctx, query, map[string]interface{}{"uuid": uuid})

	defer s.closeRows(rows)

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

func (s *EventStorage) UpdateEvent(ctx context.Context, event calendar.Event) (calendar.Event, error) {
	var updated calendar.Event

	query := `UPDATE event 
		SET (user_id, title, description, start, "end", notice_time) = 
			(:userId, :title, :description, :start, :end, :noticeTime)
		WHERE uuid = :uuid
		RETURNING uuid;`

	rows, err := s.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"uuid":        event.UUID,
		"userId":      event.UserId,
		"title":       event.Title,
		"description": event.Description,
		"start":       event.Start,
		"end":         event.End,
		"noticeTime":  event.NotifyTime,
	})

	defer s.closeRows(rows)

	if err != nil {
		return updated, err
	}
	var uuid int64

	rows.Next()
	err = rows.Scan(&uuid)

	if err != nil {
		return updated, err
	}

	updated, err = s.GetEventById(ctx, uuid)

	return updated, err
}

func (s *EventStorage) closeRows(rows *sqlx.Rows) {
	err := rows.Close()

	if err != nil {
		s.logger.Errorf("%s", err.Error())
	}
}
