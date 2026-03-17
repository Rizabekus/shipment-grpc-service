package shipment

import (
	"errors"
	"time"
)

var (
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrDuplicateStatus         = errors.New("duplicate status")
)

type Shipment struct {
	ReferenceNumber int64
	Origin          string
	Destination     string
	CurrentStatus   Status
	Details         string
	Amount          int64
	DriverRevenue   int64
	Events          []ShipmentEvent
}

func (s *Shipment) AddEvent(newStatus Status) error {
	valid := false
	for _, st := range validTransitions[s.CurrentStatus] {
		if st == newStatus {
			valid = true
			break
		}
	}
	if !valid {
		return ErrInvalidStatusTransition
	}

	event := ShipmentEvent{Status: newStatus, Timestamp: time.Now()}
	s.Events = append(s.Events, event)
	s.CurrentStatus = newStatus
	return nil
}
