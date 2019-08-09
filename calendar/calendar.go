package calendar

//go:generate protoc --go_out=. calendar.proto

import (
	uuid "github.com/satori/go.uuid"
)

type Calendar struct {
	events map[string]Event
}

func (c *Calendar) AddEvent(event Event) Event {
	u := uuid.NewV4()
	event.UUID = string(u[:uuid.Size])
	c.events[event.UUID] = event

	return event
}

func (c *Calendar) getEventById(uuid string) (e Event, ok bool) {
	event, ok := c.events[uuid]
	return event, ok
}
