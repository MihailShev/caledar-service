package main

import (
	"bytes"
	"fmt"
	"github.com/MihailShev/caledar-service/calendar"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

const (
	url         = "http://localhost:8080/"
	contentType = "application/protobuf"
)

func TestAddEvent(t *testing.T) {

	event := createEvent()
	got := postEvent(event)

	if event.UUID != uint64(0) {
		t.Error("Expected UUID is not equal 0, but got UUID:", got.UUID)
	}

	if got.Title != event.Title || got.Start.Seconds != event.Start.Seconds || got.End.Seconds != event.End.Seconds {
		t.Error("Add event - some fields is not equal expected:", *event, "got:", got)
	}
}

func TestGetEvent(t *testing.T) {
	expected := postEvent(createEvent())

	got := getEvent(expected.UUID)

	if expected.UUID != got.UUID {
		t.Error("Get event - expected:", expected, "got:", got)
	}
}

func TestUpdateEvent(t *testing.T) {
	event := postEvent(createEvent())
	expDescription := "Updated description"

	updatedEvent := updateEvent(event.UUID, &calendar.Event{Description: expDescription})

	if expDescription != updatedEvent.Description {
		t.Error("Expected updated description:", event.Description, "got:", updatedEvent.Description)
	}
}

func createEvent() *calendar.Event {
	now := time.Now()

	start := &timestamp.Timestamp{
		Seconds: now.Unix(),
	}

	end := &timestamp.Timestamp{
		Seconds: now.Add(time.Minute * 15).Unix(),
	}

	return &calendar.Event{
		Title:       "Some event",
		Start:       start,
		End:         end,
		Description: "Test event",
		UserId:      1,
		NoticeTime:  5,
	}
}

func getEvent(uuid uint64) *calendar.Event {
	resp, err := http.Get(fmt.Sprintf("%s%d", url, uuid))

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	event := &calendar.Event{}

	err = proto.Unmarshal(data, event)

	if err != nil {
		log.Fatal(err)
	}

	return event
}

func postEvent(e *calendar.Event) *calendar.Event {
	message, err := proto.Marshal(e)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(message))

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	event := &calendar.Event{}

	err = proto.Unmarshal(data, event)

	if err != nil {
		log.Fatal(err)
	}

	return event
}

func updateEvent(uuid uint64, putBody *calendar.Event) *calendar.Event {
	message, err := proto.Marshal(putBody)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s%d", url, uuid), contentType, bytes.NewBuffer(message))

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	event := &calendar.Event{}
	err = proto.Unmarshal(data, event)

	if err != nil {
		log.Fatal(err)
	}

	return event
}
