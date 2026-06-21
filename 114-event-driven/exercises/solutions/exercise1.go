package solutions

// DomainEvent represents something that happened.
type DomainEvent struct{ Name string; Payload string }

// OrderWithEvents emits events on state change.
type OrderWithEvents struct {
	ID     string
	Status string
	events []DomainEvent
}

func (o *OrderWithEvents) Ship() {
	o.Status = "shipped"
	o.events = append(o.events, DomainEvent{Name: "OrderShipped", Payload: o.ID})
}

func (o *OrderWithEvents) Events() []DomainEvent { return o.events }