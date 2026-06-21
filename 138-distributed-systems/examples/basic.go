// Package examples demonstrates Distributed Systems fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Distributed Systems example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Distributed Systems — Basic Example ===\n")
	result := demonstrateDistributedSystems()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDistributedSystems() string {
	return "Distributed Systems basic demonstration"
}
