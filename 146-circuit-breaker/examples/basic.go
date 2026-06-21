// Package examples demonstrates Circuit Breaker fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Circuit Breaker example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Circuit Breaker — Basic Example ===\n")
	result := demonstrateCircuitBreaker()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCircuitBreaker() string {
	return "Circuit Breaker basic demonstration"
}
