// Package examples demonstrates Backtracking fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Backtracking example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Backtracking — Basic Example ===\n")
	result := demonstrateBacktracking()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateBacktracking() string {
	return "Backtracking basic demonstration"
}
