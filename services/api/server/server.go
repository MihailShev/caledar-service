package main

import (
	"context"
	"github.com/MihailShev/calendar-service/pkg/config"
	"github.com/MihailShev/calendar-service/pkg/connector"
	"github.com/MihailShev/calendar-service/services/api/internal/calendar"
	"github.com/MihailShev/calendar-service/services/api/internal/db"
	"github.com/MihailShev/calendar-service/services/api/internal/grpc"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jackc/pgx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"
)

type calendarServer struct {
	service calendar.Calendar
}

type Config struct {
	DNS  string
	GRPC string
}

func main() {
	logger := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	var config = Config{}
	err := conf.Read("./", &config)

	if err != nil {
		logger.Fatal(err)
	}

	rep, err := db.NewEventStorage(logger, connector.Config{Dns: config.DNS})

	if err != nil {
		logger.Fatal(err)
	}

	cl, err := calendar.NewCalendar(rep, logger)

	if err != nil {
		logger.Fatalln(err)
	}

	server := calendarServer{service: cl}

	lis, err := net.Listen("tcp", config.GRPC)

	if err != nil {
		logger.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptorWithLogger(logger)))
	reflection.Register(grpcServer)
	calendarpb.RegisterCalendarServer(grpcServer, &server)

	logger.Infof("GRPC listen %s", config.GRPC)

	err = grpcServer.Serve(lis)

	if err != nil {
		logger.Fatal(err)
	}
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
		Start:       time.Unix(event.Start.Seconds, int64(event.Start.Nanos)),
		End:         time.Unix(event.End.Seconds, int64(event.End.Nanos)),
		NotifyTime:  time.Unix(event.NotifyTime.Seconds, int64(event.NotifyTime.Nanos)),
		Title:       event.Title,
	}
}

func mapEventToEventpb(event *calendar.Event) *calendarpb.Event {
	return &calendarpb.Event{
		UUID:        event.UUID,
		Title:       event.Title,
		Start:       &timestamp.Timestamp{Seconds: event.Start.Unix()},
		End:         &timestamp.Timestamp{Seconds: event.End.Unix()},
		NotifyTime:  &timestamp.Timestamp{Seconds: event.NotifyTime.Unix()},
		Description: event.Description,
		UserId:      event.UserId,
	}
}
