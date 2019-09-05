//go:generate protoc --proto_path ../calendarpb/ --go_out=plugins=grpc:../calendarpb calendar.proto

package main

import (
	"context"
	"fmt"
	"github.com/MihailShev/calendar-service/calendar"
	"github.com/MihailShev/calendar-service/calendarpb"
	repository "github.com/MihailShev/calendar-service/db"
	_ "github.com/jackc/pgx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
}

func main() {
	repository.Connect()
	server := calendarServer{service: calendar.NewCalendar()}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		grpclog.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(loggerInterceptor))
	reflection.Register(grpcServer)

	calendarpb.RegisterCalendarServer(grpcServer, &server)

	grpclog.Infof("GRPC listen 0.0.0.0:50051")
	err = grpcServer.Serve(lis)

	if err != nil {
		grpclog.Fatal(err)
	}
}

type calendarServer struct {
	service calendar.Calendar
}

func (s *calendarServer) CreateEvent(ctx context.Context,
	req *calendarpb.CreateEventReq) (*calendarpb.CreateEventRes, error) {

	event := mapEventpbToEvent(req.GetEvent())

	added := s.service.AddEvent(*event)

	return &calendarpb.CreateEventRes{UUID: added.UUID}, nil
}

func (s *calendarServer) GetEvent(ctx context.Context,
	req *calendarpb.GetEventReq) (*calendarpb.GetEventRes, error) {

	event, ok := s.service.GetEventByUUID(req.GetUUID())

	if !ok {
		return &calendarpb.GetEventRes{
			Event: nil,
			Error: fmt.Sprintf("Event with uuid: %d not found", req.UUID),
		}, nil
	}

	return &calendarpb.GetEventRes{
		Event: mapEventToEventpb(&event),
		Error: "",
	}, nil
}

func (s *calendarServer) UpdateEvent(ctx context.Context,
	req *calendarpb.UpdateEventReq) (*calendarpb.UpdateEventRes, error) {

	event := mapEventpbToEvent(req.GetEvent())

	updatedEvent, err := s.service.ReplaceEvent(*event)

	if err != nil {
		return &calendarpb.UpdateEventRes{
			Event: nil,
			Error: fmt.Sprintf("Event with uuid: %d not found", req.Event.UUID),
		}, nil
	}

	return &calendarpb.UpdateEventRes{
		Event: mapEventToEventpb(&updatedEvent),
		Error: "",
	}, nil
}

func mapEventpbToEvent(event *calendarpb.Event) *calendar.Event {
	return &calendar.Event{
		UUID:        event.UUID,
		UserId:      event.UserId,
		Description: event.Description,
		End:         *event.End,
		Start:       *event.Start,
		NoticeTime:  event.NoticeTime,
		Title:       event.Title,
	}
}

func mapEventToEventpb(event *calendar.Event) *calendarpb.Event {
	return &calendarpb.Event{
		UUID:        event.UUID,
		Title:       event.Title,
		NoticeTime:  event.NoticeTime,
		Start:       &event.Start,
		End:         &event.End,
		Description: event.Description,
		UserId:      event.UserId,
	}
}
