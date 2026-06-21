package solutions

import "fmt"

// Divide returns a/b with wrapped error on division by zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide %d by zero: %w", a, ErrDivZero)
	}
	return a / b, nil
}

// ErrDivZero is returned when divisor is zero.
var ErrDivZero = fmt.Errorf("division by zero")