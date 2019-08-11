package calendar

import (
	"errors"
	"fmt"
)

//go:generate protoc --go_out=. calendar.proto
var counter = uint64(0)

type Calendar struct {
	events map[uint64]Event
}

func NewCalendar() Calendar {
	c := Calendar{events: make(map[uint64]Event)}
	return c
}

func (c *Calendar) AddEvent(event Event) Event {
	counter++
	event.UUID = counter
	c.events[event.UUID] = event

	return event
}

func (c *Calendar) GetEventByUUID(uuid uint64) (e Event, ok bool) {
	event, ok := c.events[uuid]
	return event, ok
}

func (c *Calendar) ReplaceEvent(event Event) error {
	_, ok := c.events[event.UUID]

	if ok {
		c.events[event.UUID] = event
		return nil
	} else {
		return errors.New(fmt.Sprintf("Event with uuid: %d not found", event.UUID))
	}
}
