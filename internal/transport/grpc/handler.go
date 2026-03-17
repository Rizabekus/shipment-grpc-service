package grpc

import (
	"context"

	pb "github.com/Rizabekus/shipment-grpc-service/proto"

	"github.com/Rizabekus/shipment-grpc-service/internal/application"
	"github.com/Rizabekus/shipment-grpc-service/internal/domain/shipment"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ShipmentHandler struct {
	usecase *application.ShipmentUsecase
	pb.UnimplementedShipmentServiceServer
}

func NewShipmentHandler(u *application.ShipmentUsecase) *ShipmentHandler {
	return &ShipmentHandler{usecase: u}
}

func (h *ShipmentHandler) CreateShipment(ctx context.Context, req *pb.CreateShipmentRequest) (*pb.ShipmentResponse, error) {
	sh := &shipment.Shipment{
		ReferenceNumber: req.GetReferenceNumber(),
		Origin:          req.GetOrigin(),
		Destination:     req.GetDestination(),
		Details:         req.GetDetails(),
		Amount:          req.GetAmount(),
		DriverRevenue:   req.GetDriverRevenue(),
		CurrentStatus:   shipment.Pending,
		Events:          []shipment.ShipmentEvent{},
	}

	if err := h.usecase.CreateShipment(sh); err != nil {
		return nil, err
	}

	return &pb.ShipmentResponse{Shipment: mapToProtoShipment(sh)}, nil
}

func (h *ShipmentHandler) GetShipment(ctx context.Context, req *pb.GetShipmentRequest) (*pb.ShipmentResponse, error) {
	sh, err := h.usecase.GetShipment(req.GetReferenceNumber())
	if err != nil {
		return nil, err
	}

	return &pb.ShipmentResponse{Shipment: mapToProtoShipment(sh)}, nil
}

func (h *ShipmentHandler) AddShipmentEvent(ctx context.Context, req *pb.AddShipmentEventRequest) (*pb.ShipmentEventResponse, error) {
	status := shipment.Status(req.GetStatus())
	if err := h.usecase.AddEvent(req.GetReferenceNumber(), status); err != nil {
		return nil, err
	}

	sh, _ := h.usecase.GetShipment(req.GetReferenceNumber())
	lastEvent := sh.Events[len(sh.Events)-1]

	return &pb.ShipmentEventResponse{
		Event: &pb.ShipmentEvent{
			Status:    string(lastEvent.Status),
			Timestamp: timestamppb.New(lastEvent.Timestamp),
		},
	}, nil
}

func (h *ShipmentHandler) GetShipmentEvents(ctx context.Context, req *pb.GetShipmentEventsRequest) (*pb.GetShipmentEventsResponse, error) {
	events, err := h.usecase.GetEvents(req.GetReferenceNumber())
	if err != nil {
		return nil, err
	}

	res := &pb.GetShipmentEventsResponse{}
	for _, e := range events {
		res.Events = append(res.Events, &pb.ShipmentEvent{
			Status:    string(e.Status),
			Timestamp: timestamppb.New(e.Timestamp),
		})
	}

	return res, nil
}

func mapToProtoShipment(s *shipment.Shipment) *pb.Shipment {
	protoEvents := []*pb.ShipmentEvent{}
	for _, e := range s.Events {
		protoEvents = append(protoEvents, &pb.ShipmentEvent{
			Status:    string(e.Status),
			Timestamp: timestamppb.New(e.Timestamp),
		})
	}

	return &pb.Shipment{
		ReferenceNumber: s.ReferenceNumber,
		Origin:          s.Origin,
		Destination:     s.Destination,
		CurrentStatus:   string(s.CurrentStatus),
		Details:         s.Details,
		Amount:          s.Amount,
		DriverRevenue:   s.DriverRevenue,
		Events:          protoEvents,
	}
}
