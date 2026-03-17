package storage

import (
	"strconv"
	"sync"

	"github.com/Rizabekus/shipment-grpc-service/internal/domain/shipment"
)

type InMemoryEventRepo struct {
	data map[int64][]shipment.ShipmentEvent
	mu   sync.RWMutex
}

func NewInMemoryEventRepo() *InMemoryEventRepo {
	return &InMemoryEventRepo{
		data: make(map[int64][]shipment.ShipmentEvent),
	}
}

func (r *InMemoryEventRepo) Add(e *shipment.ShipmentEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[e.ShipmentID] = append(r.data[e.ShipmentID], *e)
	return nil
}

func (r *InMemoryEventRepo) List(shipmentID string) ([]shipment.ShipmentEvent, error) {
	id, _ := strconv.ParseInt(shipmentID, 10, 64)
	r.mu.RLock()
	defer r.mu.RUnlock()

	events, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	return events, nil
}
