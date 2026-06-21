// Package examples demonstrates Monotonic Stack fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Monotonic Stack example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Monotonic Stack — Basic Example ===\n")
	result := demonstrateMonotonicStack()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMonotonicStack() string {
	return "Monotonic Stack basic demonstration"
}
