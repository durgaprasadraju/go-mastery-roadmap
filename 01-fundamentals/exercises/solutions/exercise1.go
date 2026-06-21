// Package solutions contains reference implementations for Go Fundamentals exercises.
package solutions

import "errors"

var ErrInvalidTemp = errors.New("temperature below absolute zero")

// CelsiusToFahrenheit converts °C to °F. Exercise 1 solution.
func CelsiusToFahrenheit(c float64) (float64, error) {
	if c < -273.15 {
		return 0, ErrInvalidTemp
	}
	return c*9/5 + 32, nil
}

// FahrenheitToCelsius converts °F to °C.
func FahrenheitToCelsius(f float64) (float64, error) {
	if f < -459.67 {
		return 0, ErrInvalidTemp
	}
	return (f - 32) * 5 / 9, nil
}

// HTTPStatus represents an HTTP status code enum. Exercise 1 extension.
type HTTPStatus int

const (
	StatusOKExplicit      HTTPStatus = 200
	StatusCreatedExplicit HTTPStatus = 201
	StatusBadRequest      HTTPStatus = 400
	StatusNotFound        HTTPStatus = 404
	StatusInternalError   HTTPStatus = 500
)

// String returns the status text.
func (s HTTPStatus) String() string {
	switch s {
	case StatusOKExplicit:
		return "OK"
	case StatusCreatedExplicit:
		return "Created"
	case StatusBadRequest:
		return "Bad Request"
	case StatusNotFound:
		return "Not Found"
	case StatusInternalError:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}

// IsSuccess reports 2xx status codes.
func (s HTTPStatus) IsSuccess() bool {
	return s >= 200 && s < 300
}

// SumInts sums a slice of integers with error handling.
func SumInts(input []int) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("empty input")
	}
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum, nil
}
