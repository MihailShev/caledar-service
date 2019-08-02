package main

import (
	"calendar/calendar"
	"fmt"
)

func main() {
	e := calendar.Event{Name:"Событие"}
	fmt.Println(e.String())
}
