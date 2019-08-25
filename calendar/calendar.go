package calendar

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"sync"
)

//go:generate protoc --proto_path ../calendarpb/ --go_out=plugins=grpc:../calendarpb calendar.proto

type Event struct {
	UUID        uint64
	Title       string
	Start       timestamp.Timestamp
	End         timestamp.Timestamp
	Description string
	UserId      uint64
	NoticeTime  uint32
}

type Calendar struct {
	lock    sync.Mutex
	events  map[uint64]Event
	counter uint64
}

func NewCalendar() Calendar {
	c := Calendar{events: make(map[uint64]Event), counter: 0, lock: sync.Mutex{}}
	return c
}

func (c *Calendar) AddEvent(event Event) Event {
	c.lock.Lock()

	c.counter++
	event.UUID = c.counter
	c.events[event.UUID] = event

	c.lock.Unlock()

	return event
}

func (c *Calendar) GetEventByUUID(uuid uint64) (e Event, ok bool) {
	event, ok := c.events[uuid]
	return event, ok
}

func (c *Calendar) ReplaceEvent(event Event) error {
	_, ok := c.events[event.UUID]

	if ok {
		c.lock.Lock()

		c.events[event.UUID] = event

		c.lock.Unlock()
		return nil
	} else {
		return errors.New(fmt.Sprintf("Event with uuid: %d not found", event.UUID))
	}
}
