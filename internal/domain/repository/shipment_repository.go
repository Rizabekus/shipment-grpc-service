package repository

import "github.com/Rizabekus/shipment-grpc-service/internal/domain/shipment"

type ShipmentRepository interface {
	Save(shipment *shipment.Shipment) error
	GetByReference(ref int64) (*shipment.Shipment, error)
}
