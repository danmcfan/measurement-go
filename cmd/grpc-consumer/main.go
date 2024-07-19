package main

import (
	"context"
	"log"
	"net"

	internal "measurement/internal/grpc"

	"google.golang.org/grpc"
)

type server struct {
	internal.UnimplementedMeasurementServiceServer
}

func (s *server) SendMeasurement(ctx context.Context, measurement *internal.Measurement) (*internal.MeasurementResponse, error) {
	return &internal.MeasurementResponse{
		Success: true,
		Message: "Measurement received successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	internal.RegisterMeasurementServiceServer(s, &server{})
	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
