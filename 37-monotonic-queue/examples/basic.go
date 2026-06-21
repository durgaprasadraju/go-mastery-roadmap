// Package examples demonstrates Monotonic Queue fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Monotonic Queue example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Monotonic Queue — Basic Example ===\n")
	result := demonstrateMonotonicQueue()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMonotonicQueue() string {
	return "Monotonic Queue basic demonstration"
}
