package calendar

import (
	"context"
	"time"
)

type Event struct {
	UUID        int64
	Title       string
	Start       time.Time
	End         time.Time
	NotifyTime  time.Time
	Description string
	UserId      uint64
}

type EventStorage interface {
	CreateEvent(ctx context.Context, event Event) (int64, error)
	GetEventById(ctx context.Context, uuid int64) (Event, error)
	UpdateEvent(ctx context.Context, event Event) (Event, error)
}

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}

type Calendar struct {
	EventStorage EventStorage
	logger       Logger
}

func NewCalendar(repository EventStorage, logger Logger) (Calendar, error) {

	c := Calendar{
		EventStorage: repository,
		logger:       logger,
	}

	return c, nil
}

func (c *Calendar) AddEvent(ctx context.Context, e Event) (int64, error) {
	id, err := c.EventStorage.CreateEvent(ctx, e)

	return id, err
}

func (c *Calendar) GetEventByUUID(ctx context.Context, uuid int64) (Event, error) {

	event, err := c.EventStorage.GetEventById(ctx, uuid)

	return event, err
}

func (c *Calendar) UpdateEvent(ctx context.Context, event Event) (Event, error) {
	event, err := c.EventStorage.UpdateEvent(ctx, event)

	return event, err
}
