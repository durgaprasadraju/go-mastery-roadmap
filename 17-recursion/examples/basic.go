// Package examples demonstrates Recursion fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Recursion example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Recursion — Basic Example ===\n")
	result := demonstrateRecursion()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRecursion() string {
	return "Recursion basic demonstration"
}
