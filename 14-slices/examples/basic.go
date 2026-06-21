// Package examples demonstrates Slices fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Slices example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Slices — Basic Example ===\n")
	result := demonstrateSlices()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSlices() string {
	return "Slices basic demonstration"
}
