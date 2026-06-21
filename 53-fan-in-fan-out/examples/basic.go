// Package examples demonstrates Fan-In / Fan-Out fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Fan-In / Fan-Out example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Fan-In / Fan-Out — Basic Example ===\n")
	result := demonstrateFanInFanOut()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFanInFanOut() string {
	return "Fan-In / Fan-Out basic demonstration"
}
