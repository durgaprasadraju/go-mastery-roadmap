// Package examples demonstrates Greedy Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Greedy Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Greedy Algorithms — Basic Example ===\n")
	result := demonstrateGreedy()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGreedy() string {
	return "Greedy Algorithms basic demonstration"
}
