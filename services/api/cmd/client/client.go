package main

import (
	"context"
	"fmt"
	"github.com/MihailShev/calendar-service/pkg/config"
	"github.com/MihailShev/calendar-service/services/api/internal/grpc"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Config struct {
	GRPC string
}

func main() {
	var config = Config{}
	err := conf.Read("../../", &config)

	if err != nil {
		log.Fatal(err)
	}

	cc, err := grpc.Dial(config.GRPC, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer func() {
		err = cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := calendarpb.NewCalendarClient(cc)

	uuid := createEvent(client)
	event := getEvent(client, uuid)

	fmt.Printf("Created event: %+v\n", event)

	event = updateEvent(client, event)

	fmt.Printf("Updated event description: %+v\n", event)
}

func createEvent(client calendarpb.CalendarClient) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	start := time.Now().Add(5 * time.Minute)
	end := start.Add(5 * time.Minute)
	notifyTime := start.Add(-5 * time.Minute)

	res, err := client.CreateEvent(ctx, &calendarpb.CreateEventReq{
		Event: &calendarpb.Event{
			Title:       "Test event",
			Start:       &timestamp.Timestamp{Seconds: start.Unix(), Nanos: 0},
			End:         &timestamp.Timestamp{Seconds: end.Unix(), Nanos: 0},
			NotifyTime:  &timestamp.Timestamp{Seconds: notifyTime.Unix(), Nanos: 0},
			Description: "test",
			UserId:      1,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	if res.Error != "" {
		log.Fatal(res.Error)
	}

	return res.GetUUID()
}

func getEvent(client calendarpb.CalendarClient, uuid int64) *calendarpb.Event {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	res, err := client.GetEvent(ctx, &calendarpb.GetEventReq{UUID: uuid})

	if err != nil {
		log.Fatal(err)
	}

	if res.Error != "" {
		log.Fatal(res)
	}

	return res.GetEvent()
}

func updateEvent(client calendarpb.CalendarClient, event *calendarpb.Event) *calendarpb.Event {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	res, err := client.UpdateEvent(ctx, &calendarpb.UpdateEventReq{
		Event: &calendarpb.Event{
			UUID:        event.UUID,
			Title:       event.Title,
			Start:       event.Start,
			End:         event.End,
			Description: "New description",
			UserId:      event.UserId,
			NotifyTime:  event.NotifyTime,
		},
	})

	if err != nil {
		log.Fatal(event)
	}

	if res.Error != "" {
		log.Fatal(res.Error)
	}

	return res.GetEvent()
}
