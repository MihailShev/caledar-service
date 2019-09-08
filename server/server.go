//go:generate protoc --proto_path ../calendarpb/ --go_out=plugins=grpc:../calendarpb calendar.proto

package main

import (
	"context"
	"github.com/MihailShev/calendar-service/calendar"
	"github.com/MihailShev/calendar-service/calendarpb"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jackc/pgx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"
)

func main() {
	logger := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)

	cal, err := calendar.NewCalendar(logger)

	if err != nil {
		logger.Fatalln(err)
	}

	server := calendarServer{service: cal}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		logger.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptorWithLogger(logger)))
	reflection.Register(grpcServer)
	calendarpb.RegisterCalendarServer(grpcServer, &server)

	logger.Infof("GRPC listen 0.0.0.0:50051")

	err = grpcServer.Serve(lis)

	if err != nil {
		logger.Fatal(err)
	}
}

type calendarServer struct {
	service calendar.Calendar
}

func (s *calendarServer) CreateEvent(ctx context.Context,
	req *calendarpb.CreateEventReq) (*calendarpb.CreateEventRes, error) {

	event := mapEventpbToEvent(req.GetEvent())

	id, err := s.service.AddEvent(ctx, *event)

	if err != nil {
		return &calendarpb.CreateEventRes{Error: err.Error()}, nil
	}

	return &calendarpb.CreateEventRes{UUID: id}, nil
}

func (s *calendarServer) GetEvent(ctx context.Context,
	req *calendarpb.GetEventReq) (*calendarpb.GetEventRes, error) {

	event, err := s.service.GetEventByUUID(ctx, req.GetUUID())

	if err != nil {
		return &calendarpb.GetEventRes{
			Event: nil,
			Error: err.Error(),
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

	updatedEvent, err := s.service.UpdateEvent(ctx, *event)

	if err != nil {
		return &calendarpb.UpdateEventRes{
			Event: nil,
			Error: err.Error(),
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
		End:         time.Unix(event.End.Seconds, int64(event.End.Nanos)),
		Start:       time.Unix(event.Start.Seconds, int64(event.Start.Nanos)),
		NoticeTime:  event.NoticeTime,
		Title:       event.Title,
	}
}

func mapEventToEventpb(event *calendar.Event) *calendarpb.Event {
	return &calendarpb.Event{
		UUID:        event.UUID,
		Title:       event.Title,
		NoticeTime:  event.NoticeTime,
		Start:       &timestamp.Timestamp{Seconds: event.Start.Unix()},
		End:         &timestamp.Timestamp{Seconds: event.Start.Unix()},
		Description: event.Description,
		UserId:      event.UserId,
	}
}
