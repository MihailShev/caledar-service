package calendar

import (
	"context"
	repository "github.com/MihailShev/calendar-service/internal/db"
	"time"
)

type Event struct {
	UUID         int64
	Title        string
	Start        time.Time
	End          time.Time
	Description  string
	UserId       uint64
	NotifyBefore uint32
}

type Logger = repository.Logger

type Calendar struct {
	repository repository.Repository
	logger     repository.Logger
}

func NewCalendar(logger Logger) (Calendar, error) {
	rep, err := repository.NewRepository(logger)

	if err != nil {
		return Calendar{}, err
	}

	c := Calendar{
		repository: rep,
		logger:     logger,
	}

	return c, nil
}

func (c *Calendar) AddEvent(ctx context.Context, e Event) (int64, error) {
	id, err := c.repository.CreateEvent(ctx, e)

	return id, err
}

func (c *Calendar) GetEventByUUID(ctx context.Context, uuid int64) (Event, error) {

	event, err := c.repository.GetEventById(ctx, uuid)

	return event, err
}

func (c *Calendar) UpdateEvent(ctx context.Context, event Event) (Event, error) {
	event, err := c.repository.UpdateEvent(ctx, event)

	return event, err
}
