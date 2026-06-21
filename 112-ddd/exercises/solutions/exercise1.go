package solutions

import "errors"

var ErrInvalidOrder = errors.New("invalid order")

// Order is a DDD aggregate root.
type Order struct {
	ID     string
	Items  []string
	Status string
}

// NewOrder creates a pending order with at least one item.
func NewOrder(id string, items []string) (*Order, error) {
	if id == "" || len(items) == 0 {
		return nil, ErrInvalidOrder
	}
	return &Order{ID: id, Items: append([]string(nil), items...), Status: "pending"}, nil
}

// Ship transitions order to shipped if pending.
func (o *Order) Ship() error {
	if o.Status != "pending" {
		return ErrInvalidOrder
	}
	o.Status = "shipped"
	return nil
}