// Package examples demonstrates Query Optimization fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Query Optimization example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Query Optimization — Basic Example ===\n")
	result := demonstrateQueryOptimization()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateQueryOptimization() string {
	return "Query Optimization basic demonstration"
}
