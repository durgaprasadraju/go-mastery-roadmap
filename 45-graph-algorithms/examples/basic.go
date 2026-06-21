// Package examples demonstrates Graph Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Graph Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Graph Algorithms — Basic Example ===\n")
	result := demonstrateGraphAlgorithms()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGraphAlgorithms() string {
	return "Graph Algorithms basic demonstration"
}
