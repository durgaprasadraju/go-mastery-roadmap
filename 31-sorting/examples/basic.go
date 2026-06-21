// Package examples demonstrates Sorting Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Sorting Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Sorting Algorithms — Basic Example ===\n")
	result := demonstrateSorting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSorting() string {
	return "Sorting Algorithms basic demonstration"
}
