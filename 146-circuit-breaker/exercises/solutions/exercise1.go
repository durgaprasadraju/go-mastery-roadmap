package solutions

import "errors"

var ErrCircuitOpen = errors.New("circuit open")

// CircuitBreaker has closed/open states (simplified).
type CircuitBreaker struct {
	failures, threshold int
	open                bool
}

func NewCircuitBreaker(threshold int) *CircuitBreaker {
	return &CircuitBreaker{threshold: threshold}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.open {
		return ErrCircuitOpen
	}
	if err := fn(); err != nil {
		cb.failures++
		if cb.failures >= cb.threshold {
			cb.open = true
		}
		return err
	}
	cb.failures = 0
	return nil
}