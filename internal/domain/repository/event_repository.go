package repository

import "github.com/Rizabekus/shipment-grpc-service/internal/domain/shipment"

type EventRepository interface {
	Add(event *shipment.ShipmentEvent) error
	List(shipmentID string) ([]shipment.ShipmentEvent, error)
}
