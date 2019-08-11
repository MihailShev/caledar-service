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

func TestAdd(t *testing.T) {

	event := createEvent()
	got := calendar.Event{}

	message, err := proto.Marshal(event)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s%s", url, "add"), contentType, bytes.NewBuffer(message))

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = proto.Unmarshal(data, &got)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(got.End, event.End, got.End.Seconds == event.End.Seconds)

	if event.UUID != uint64(0) {
		t.Error("Expected UUID is not equal 0, but got UUID:", got.UUID)
	}

	if got.Title != event.Title || got.Start.Seconds != event.Start.Seconds || got.End.Seconds != event.End.Seconds {
		t.Error("Add event, some fields is not equal expected:", *event, "got:", got)
	}
}

func TestGetEvent(t *testing.T) {

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
