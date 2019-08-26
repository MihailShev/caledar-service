package main

import (
	"context"
	"github.com/MihailShev/calendar-service/calendarpb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	client := calendarpb.NewCalendarClient(cc)

	start := time.Now()
	end := start.Add(5 * time.Minute)

	res, err := client.CreateEvent(ctx, &calendarpb.CreateEventReq{
		Event: &calendarpb.Event{
			Title:       "Some event",
			Start:       &timestamp.Timestamp{Seconds: start.Unix(), Nanos: 0},
			End:         &timestamp.Timestamp{Seconds: end.Unix(), Nanos: 0},
			Description: "some description",
			NoticeTime:  5,
			UserId:      1,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

}
