// Package solutions contains reference implementations for Fast & Slow Pointers exercises.
package solutions

import "errors"

var ErrInvalidInput = errors.New("fast-slow-pointers: invalid input")

// Exercise1Core demonstrates the fundamental Fast & Slow Pointers pattern.
// Time: O(n) typical | Space: O(1) auxiliary for this demo.
func Exercise1Core(input []int) (int, error) {
	if len(input) == 0 {
		return 0, ErrInvalidInput
	}
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum, nil
}

// Exercise1Transform applies a Fast & Slow Pointers-specific transformation.
func Exercise1Transform(input string) string {
	if input == "" {
		return input
	}
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
