package shipment

type Status string

const (
	Pending   Status = "pending"
	PickedUp  Status = "picked_up"
	InTransit Status = "in_transit"
	Delivered Status = "delivered"
	Completed Status = "completed"
)

var validTransitions = map[Status][]Status{
	Pending:   {PickedUp},
	PickedUp:  {InTransit},
	InTransit: {Delivered},
	Delivered: {Completed},
	Completed: {},
}
