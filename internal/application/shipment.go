package application

import (
	"time"

	"github.com/Rizabekus/shipment-grpc-service/internal/domain/repository"
	"github.com/Rizabekus/shipment-grpc-service/internal/domain/shipment"
)

type ShipmentUsecase struct {
	shipmentRepo repository.ShipmentRepository
	eventRepo    repository.EventRepository
}

func NewShipmentUsecase(shipmentRepo repository.ShipmentRepository, eventRepo repository.EventRepository) *ShipmentUsecase {
	return &ShipmentUsecase{
		shipmentRepo: shipmentRepo,
		eventRepo:    eventRepo,
	}
}

func (u *ShipmentUsecase) CreateShipment(s *shipment.Shipment) error {

	s.CurrentStatus = shipment.Pending
	event := shipment.ShipmentEvent{
		Status:    shipment.Pending,
		Timestamp: time.Now(),
	}
	s.Events = []shipment.ShipmentEvent{event}

	if err := u.shipmentRepo.Save(s); err != nil {
		return err
	}

	return u.eventRepo.Add(&event)
}

func (u *ShipmentUsecase) AddEvent(ref int64, newStatus shipment.Status) error {
	s, err := u.shipmentRepo.Get(ref)
	if err != nil {
		return err
	}

	if err := s.AddEvent(newStatus); err != nil {
		return err
	}

	lastEvent := &s.Events[len(s.Events)-1]

	if err := u.eventRepo.Add(lastEvent); err != nil {
		return err
	}
	return u.shipmentRepo.Save(s)
}

func (u *ShipmentUsecase) GetShipment(ref int64) (*shipment.Shipment, error) {
	return u.shipmentRepo.Get(ref)
}

func (u *ShipmentUsecase) GetEvents(ref int64) ([]shipment.ShipmentEvent, error) {
	return u.eventRepo.List(string(ref))
}
