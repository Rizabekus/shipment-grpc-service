package main

import (
	"log"
	"net"
	"os"

	"github.com/Rizabekus/shipment-grpc-service/internal/application"
	"github.com/Rizabekus/shipment-grpc-service/internal/infrastructure/storage"
	handler "github.com/Rizabekus/shipment-grpc-service/internal/transport/grpc"
	pb "github.com/Rizabekus/shipment-grpc-service/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	shipmentRepo := storage.NewInMemoryShipmentRepo()
	eventRepo := storage.NewInMemoryEventRepo()

	usecase := application.NewShipmentUsecase(shipmentRepo, eventRepo)

	shipmentHandler := handler.NewShipmentHandler(usecase)

	grpcServer := grpc.NewServer()

	pb.RegisterShipmentServiceServer(grpcServer, shipmentHandler)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
