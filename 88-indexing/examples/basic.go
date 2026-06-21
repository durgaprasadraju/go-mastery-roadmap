// Package examples demonstrates Indexing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Indexing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Indexing — Basic Example ===\n")
	result := demonstrateIndexing()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateIndexing() string {
	return "Indexing basic demonstration"
}
