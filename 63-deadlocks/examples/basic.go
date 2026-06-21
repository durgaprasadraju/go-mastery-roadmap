// Package examples demonstrates Deadlocks fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Deadlocks example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Deadlocks — Basic Example ===\n")
	result := demonstrateDeadlocks()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDeadlocks() string {
	return "Deadlocks basic demonstration"
}
