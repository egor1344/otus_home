package main

import (
	"context"
	"log"
	"net"
	db "otus_home/calendar/db"
	calendar_server "otus_home/calendar/models/server"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) AddEvent(ctx context.Context, in *calendar_server.AddEventRequest) (*calendar_server.AddEventResponse, error) {
	log.Println("add event", in.GetEvent())
	err := db.AddEvent(in.GetEvent())
	if err != nil {
		return nil, err
	}
	response := calendar_server.AddEventResponse{
		Status: "True",
	}
	return &response, nil

}

func (s *server) UpdateEvent(ctx context.Context, in *calendar_server.UpdateEventRequest) (*calendar_server.UpdateEventResponse, error) {
	log.Println("update event")
	response := calendar_server.UpdateEventResponse{
		Status: "True",
	}
	return &response, nil
}
func (s *server) DeleteEvent(ctx context.Context, in *calendar_server.DeleteEventRequest) (*calendar_server.DeleteEventResponse, error) {
	log.Println("delete event")
	response := calendar_server.DeleteEventResponse{
		Status: "True",
	}
	return &response, nil
}

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:3000")
	log.Println("server run in 127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
	gs := grpc.NewServer()
	calendar_server.RegisterCalendarEventServer(gs, &server{})
	err = gs.Serve(conn)
	if err != nil {
		log.Fatal(err)
	}
}
