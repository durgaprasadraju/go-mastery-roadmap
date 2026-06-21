// Package examples demonstrates Concurrency Patterns fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Concurrency Patterns example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Concurrency Patterns — Basic Example ===\n")
	result := demonstrateConcurrencyPatterns()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateConcurrencyPatterns() string {
	return "Concurrency Patterns basic demonstration"
}
