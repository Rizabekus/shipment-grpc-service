package shipment

type Status string

const (
	Pending   Status = "PENDING"
	PickedUp  Status = "PICKED_UP"
	InTransit Status = "IN_TRANSIT"
	Delivered Status = "DELIVERED"
	Completed Status = "COMPLETED"
)

var validTransitions = map[Status][]Status{
	Pending:   {PickedUp},
	PickedUp:  {InTransit},
	InTransit: {Delivered},
	Delivered: {Completed},
	Completed: {},
}
