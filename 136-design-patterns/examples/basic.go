// Package examples demonstrates Design Patterns fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Design Patterns example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Design Patterns — Basic Example ===\n")
	result := demonstrateDesignPatterns()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDesignPatterns() string {
	return "Design Patterns basic demonstration"
}
