package shipment

import "time"

type ShipmentEvent struct {
	Status    Status
	Timestamp time.Time
}
