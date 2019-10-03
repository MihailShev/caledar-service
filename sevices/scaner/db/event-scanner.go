package db

import (
	"github.com/MihailShev/calendar-service/config"
	"github.com/MihailShev/calendar-service/pkg/connector"
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

type EventScanner struct {
	db *sqlx.DB
}

func NewEventScanner(config config.Config) (*EventScanner, error) {
	db, err := connector.Connect(config.Dns)

	if err != nil {
		return &EventScanner{}, err
	}

	return &EventScanner{db: db}, err
}

func (s *EventScanner) Scan(dateFrom time.Time, dateTo time.Time) ([]Event, error) {
	eventList := make([]Event, 0)
	query := `SELECT * FROM EVENT WHERE notice_time between :dateFrom AND :dateTo`

	rows, err := s.db.NamedQuery(query, map[string]interface{}{
		"dateFrom": dateFrom,
		"dateTo":   dateTo,
	})

	if err != nil {
		return eventList, err
	}

	for rows.Next() {
		var event Event
		err = rows.StructScan(&event)

		if err != nil {
			return eventList, err
		}

		eventList = append(eventList, event)
	}

	err = rows.Close()

	return eventList, err
}
