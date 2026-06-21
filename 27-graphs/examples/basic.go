// Package examples demonstrates Graphs fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Graphs example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Graphs — Basic Example ===\n")
	result := demonstrateGraphs()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGraphs() string {
	return "Graphs basic demonstration"
}
