package storage

import (
	"errors"
	"sync"

	"github.com/Rizabekus/shipment-grpc-service/internal/domain/shipment"
)

type InMemoryShipmentRepo struct {
	data map[int64]*shipment.Shipment
	mu   sync.RWMutex
}

func NewInMemoryShipmentRepo() *InMemoryShipmentRepo {
	return &InMemoryShipmentRepo{
		data: make(map[int64]*shipment.Shipment),
	}
}

func (r *InMemoryShipmentRepo) Save(s *shipment.Shipment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[s.ReferenceNumber] = s
	return nil
}

func (r *InMemoryShipmentRepo) Get(id int64) (*shipment.Shipment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.data[id]
	if !ok {
		return nil, errors.New("shipment not found")
	}
	return s, nil
}
