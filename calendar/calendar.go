package calendar

//go:generate protoc --go_out=. calendar.proto

type Calendar struct {
	events []Event
}

func (c *Calendar) addEvent (event Event) {
	c.events = append(c.events, event)
}