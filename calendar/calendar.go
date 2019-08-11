package calendar

//go:generate protoc --go_out=. calendar.proto
var counter = uint64(0)

type Calendar struct {
	events map[uint64]Event
}

func (c *Calendar) AddEvent(event Event) Event {
	counter++
	event.UUID = counter
	c.events[event.UUID] = event

	return event
}

func (c *Calendar) getEventById(uuid uint64) (e Event, ok bool) {
	event, ok := c.events[uuid]
	return event, ok
}

func NewCalendar() Calendar {
	c := Calendar{events: make(map[uint64]Event)}
	return c
}
