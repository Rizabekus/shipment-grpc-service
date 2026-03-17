package shipment

import "testing"

func TestShipment_AddEvent_TableDriven(t *testing.T) {
	cases := []struct {
		name            string
		initialStatus   Status
		newStatus       Status
		wantErr         error
		wantFinalStatus Status
	}{
		{"Pending → PickedUp", Pending, PickedUp, nil, PickedUp},
		{"PickedUp → InTransit", PickedUp, InTransit, nil, InTransit},
		{"InTransit → Delivered", InTransit, Delivered, nil, Delivered},
		{"Delivered → Completed", Delivered, Completed, nil, Completed},
		{"Pending → Delivered (invalid)", Pending, Delivered, ErrInvalidStatusTransition, Pending},
		{"InTransit → PickedUp (invalid)", InTransit, PickedUp, ErrInvalidStatusTransition, InTransit},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := &Shipment{
				ReferenceNumber: 1,
				Origin:          "A",
				Destination:     "B",
				CurrentStatus:   c.initialStatus,
				Events:          []ShipmentEvent{},
			}

			err := s.AddEvent(c.newStatus)
			if err != c.wantErr {
				t.Errorf("expected error %v, got %v", c.wantErr, err)
			}

			if s.CurrentStatus != c.wantFinalStatus {
				t.Errorf("expected final status %s, got %s", c.wantFinalStatus, s.CurrentStatus)
			}

			expectedEvents := 0
			if c.wantErr == nil {
				expectedEvents = 1
			}
			if len(s.Events) != expectedEvents {
				t.Errorf("expected %d events, got %d", expectedEvents, len(s.Events))
			}
		})
	}
}
