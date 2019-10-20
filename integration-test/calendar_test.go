package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/MihailShev/calendar-service/pkg/grpc"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"time"
)

type CreateEventParams struct {
	Title         string
	Description   string
	StartAfterNow time.Duration
	Duration      time.Duration
	NotifyTime    time.Duration
	UserId        uint64
}

type GetEventParams struct {
	Title       string
	Description string
	UserId      uint64
}

type CalendarTest struct {
	createEventRes calendarpb.CreateEventRes
	getEventRes    calendarpb.GetEventRes
	updateEventRes calendarpb.UpdateEventRes
	client         calendarpb.CalendarClient
	amqpConn       *amqp.Connection
	ampqCh         *amqp.Channel
	notifyMessage  []byte
	stopSignal     chan struct{}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (test *CalendarTest) beforeFeature(f *gherkin.Feature) {
	test.stopSignal = make(chan struct{})

	var err error
	test.amqpConn, err = amqp.Dial("amqp://guest:guest@queue:5672/")
	panicOnErr(err)

	test.ampqCh, err = test.amqpConn.Channel()
	panicOnErr(err)

	q, err := test.ampqCh.QueueDeclare(
		"notify_test", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	panicOnErr(err)

	err = test.ampqCh.QueueBind("notify_test", "", "notifyExchange", false, nil)
	panicOnErr(err)

	msgs, err := test.ampqCh.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	panicOnErr(err)

	go func(stop chan struct{}) {
		for {
			select {
			case <-stop:
				return
			case mess := <-msgs:
				test.notifyMessage = mess.Body
			}
		}
	}(test.stopSignal)
}

func (test *CalendarTest) afterFeature(feature *gherkin.Feature) {
	test.stopSignal <- struct{}{}

	err := test.ampqCh.Close()
	panicOnErr(err)

	err = test.amqpConn.Close()
	panicOnErr(err)
}

func (test *CalendarTest) iCreateCalendarClient() error {
	cc, err := grpc.Dial("server:50051", grpc.WithInsecure())

	if err != nil {
		return err
	}

	test.client = calendarpb.NewCalendarClient(cc)

	return err
}

func (test *CalendarTest) iSendMessageCreateEventWithData(data *gherkin.DocString) error {
	var eventParams = CreateEventParams{}
	err := json.Unmarshal([]byte(data.Content), &eventParams)

	if err != nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	start := time.Now().Add(eventParams.StartAfterNow * time.Minute)
	end := start.Add(eventParams.Duration * time.Minute)
	notifyTime := start.Add(eventParams.NotifyTime * time.Minute)

	res, err := test.client.CreateEvent(ctx, &calendarpb.CreateEventReq{
		Event: &calendarpb.Event{
			Title:       eventParams.Title,
			Start:       &timestamp.Timestamp{Seconds: start.Unix(), Nanos: 0},
			End:         &timestamp.Timestamp{Seconds: end.Unix(), Nanos: 0},
			NotifyTime:  &timestamp.Timestamp{Seconds: notifyTime.Unix(), Nanos: 0},
			Description: eventParams.Description,
			UserId:      eventParams.UserId,
		},
	})

	if err != nil {
		return err
	}

	test.createEventRes = *res

	return err
}

func (test *CalendarTest) theResponseErrorShouldBeEmpty() error {
	if test.createEventRes.Error != "" {
		return fmt.Errorf(test.createEventRes.Error)
	}

	return nil
}

func (test *CalendarTest) iSendMessageGetEvent() error {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	res, err := test.client.GetEvent(ctx, &calendarpb.GetEventReq{UUID: test.createEventRes.GetUUID()})

	if err != nil {
		return err
	}

	if res.Error != "" {
		return fmt.Errorf(res.Error)
	}

	test.getEventRes = *res

	return nil
}

func (test *CalendarTest) iReceiveEventWithParams(params *gherkin.DocString) error {
	var eventParams = GetEventParams{}
	err := json.Unmarshal([]byte(params.Content), &eventParams)

	if err != nil {
		return err
	}

	event := test.getEventRes.GetEvent()

	if eventParams.UserId != event.UserId {
		return fmt.Errorf("expected userId %d, got %d", eventParams.UserId, event.UserId)
	}

	if eventParams.Description != event.Description {
		return fmt.Errorf("expected description %s, got %s", eventParams.Description, event.Description)
	}

	if eventParams.Title != event.Title {
		return fmt.Errorf("expected title %s, got %s", eventParams.Title, event.Title)
	}

	return nil
}

func (test *CalendarTest) iSendMessageUpdateEventWithNewTitle(newTitle string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	res, err := test.client.UpdateEvent(ctx, &calendarpb.UpdateEventReq{
		Event: &calendarpb.Event{
			UUID:        test.getEventRes.Event.UUID,
			Title:       newTitle,
			Start:       test.getEventRes.Event.Start,
			End:         test.getEventRes.Event.End,
			Description: test.getEventRes.Event.Description,
			UserId:      test.getEventRes.Event.UserId,
			NotifyTime:  test.getEventRes.Event.NotifyTime,
		},
	})

	if err != nil {
		return err
	}

	if res.Error != "" {
		return fmt.Errorf(res.Error)
	}

	test.updateEventRes = *res

	return nil
}

func (test *CalendarTest) eventTitleMatch(newTitle string) error {
	if newTitle != test.updateEventRes.GetEvent().Title {
		return fmt.Errorf("expected title %s, got %s", newTitle, test.updateEventRes.GetEvent().Title)
	}

	return nil
}

func (test *CalendarTest) iReceivedNotifyMessageWithCreatedEvent() error {
	fmt.Println("Waiting 10 seconds for the scanner publish notify message.")
	wait(10)

	var event = struct {
		UUID int64
	}{}

	err := json.Unmarshal(test.notifyMessage, &event)

	if err != nil {
		return err
	}

	if event.UUID != test.createEventRes.GetUUID() {
		return fmt.Errorf("expected userId %d, got %d", test.createEventRes.GetUUID(), event.UUID)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	test := CalendarTest{}
	s.BeforeFeature(test.beforeFeature)

	//  Scenario: Create event
	s.Step(`^I create calendar client$`, test.iCreateCalendarClient)
	s.Step(`^I send message create event with params:$`, test.iSendMessageCreateEventWithData)
	s.Step(`^The response error should be empty$`, test.theResponseErrorShouldBeEmpty)
	s.Step(`^I received notify message with created event$`, test.iReceivedNotifyMessageWithCreatedEvent)

	// Scenario Get event
	s.Step(`^I send message get event$`, test.iSendMessageGetEvent)
	s.Step(`^I receive event with params:$`, test.iReceiveEventWithParams)

	// Scenario Update event title
	s.Step(`^I send message update event with new title "([^"]*)"$`, test.iSendMessageUpdateEventWithNewTitle)
	s.Step(`^Event title match "([^"]*)"$`, test.eventTitleMatch)

	s.AfterFeature(test.afterFeature)
}
