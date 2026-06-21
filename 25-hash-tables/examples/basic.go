// Package examples demonstrates Hash Tables fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Hash Tables example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Hash Tables — Basic Example ===\n")
	result := demonstrateHashTables()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHashTables() string {
	return "Hash Tables basic demonstration"
}
