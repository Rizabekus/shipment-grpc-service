package shipment

import "time"

type ShipmentEvent struct {
	ShipmentID int64
	Status     Status
	Timestamp  time.Time
}
