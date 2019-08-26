package main

import (
	"context"
	"fmt"
	"github.com/MihailShev/calendar-service/calendar"
	"github.com/MihailShev/calendar-service/calendarpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type calendarServer struct {
	service calendar.Calendar
}

func (s *calendarServer) CreateEvent(ctx context.Context, req *calendarpb.CreateEventReq) (*calendarpb.CreateEventRes, error) {
	event := req.GetEvent()

	added := s.service.AddEvent(calendar.Event{
		UserId:      event.UserId,
		Description: event.Description,
		End:         *event.End,
		Start:       *event.Start,
		NoticeTime:  event.NoticeTime,
		Title:       event.Title,
	})

	return &calendarpb.CreateEventRes{UUID: added.UUID}, nil
}

func (s *calendarServer) GetEvent(ctx context.Context, req *calendarpb.GetEventReq) (*calendarpb.GetEventRes, error) {
	event, ok := s.service.GetEventByUUID(req.GetUUID())

	if !ok {
		return &calendarpb.GetEventRes{
			Event: nil,
			Error: fmt.Sprintf("Event with uuid: %d not found", req.UUID),
		}, nil
	}

	return &calendarpb.GetEventRes{
		Event: &calendarpb.Event{
			UUID:        event.UUID,
			Title:       event.Title,
			NoticeTime:  event.NoticeTime,
			Start:       &event.Start,
			End:         &event.End,
			Description: event.Description,
			UserId:      event.UserId,
		},
		Error: "",
	}, nil

}

func (s *calendarServer) UpdateEvent(ctx context.Context, req *calendarpb.UpdateEventReq) (*calendarpb.UpdateEventRes, error) {
	event := req.GetEvent()

	updatedEvent, err := s.service.ReplaceEvent(calendar.Event{
		UserId:      event.UserId,
		Description: event.Description,
		End:         *event.End,
		Start:       *event.Start,
		NoticeTime:  event.NoticeTime,
		Title:       event.Title,
	})

	if err != nil {
		return &calendarpb.UpdateEventRes{
			Event: nil,
			Error: fmt.Sprintf("Event with uuid: %d not found", req.Event.UUID),
		}, nil
	}

	return &calendarpb.UpdateEventRes{
		Event: &calendarpb.Event{
			UUID:        updatedEvent.UUID,
			Title:       updatedEvent.Title,
			NoticeTime:  updatedEvent.NoticeTime,
			Start:       &updatedEvent.Start,
			End:         &updatedEvent.End,
			Description: updatedEvent.Description,
			UserId:      updatedEvent.UserId,
		},
		Error: "",
	}, nil
}

func main() {
	server := calendarServer{service: calendar.NewCalendar()}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	calendarpb.RegisterCalendarServer(grpcServer, &server)
	grpcServer.Serve(lis)
}
