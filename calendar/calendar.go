package calendar

import (
	"errors"
	"fmt"
	repository "github.com/MihailShev/calendar-service/db"
	"sync"
)

type Event = repository.EventModel

type Calendar struct {
	lock       sync.Mutex
	events     map[int64]Event
	counter    int64
	repository repository.Repository
}

func NewCalendar() (Calendar, error) {
	rep, err := repository.NewRepository()

	if err != nil {
		return Calendar{}, err
	}

	c := Calendar{
		events:     make(map[int64]Event),
		counter:    0,
		lock:       sync.Mutex{},
		repository: rep,
	}

	return c, nil
}

func (c *Calendar) AddEvent(e Event) (int64, error) {
	id, err := c.repository.CreateEvent(e)

	return id, err
}

func (c *Calendar) GetEventByUUID(uuid int64) (Event, error) {

	event, err := c.repository.GetEventById(uuid)

	return event, err
}

func (c *Calendar) ReplaceEvent(event Event) (Event, error) {
	_, ok := c.events[event.UUID]

	if ok {
		c.lock.Lock()

		c.events[event.UUID] = event

		c.lock.Unlock()
		return event, nil
	} else {
		return event, errors.New(fmt.Sprintf("EventModel with uuid: %d not found", event.UUID))
	}
}
