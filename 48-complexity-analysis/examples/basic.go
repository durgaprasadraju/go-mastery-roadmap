// Package examples demonstrates Complexity Analysis fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Complexity Analysis example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Complexity Analysis — Basic Example ===\n")
	result := demonstrateComplexityAnalysis()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateComplexityAnalysis() string {
	return "Complexity Analysis basic demonstration"
}
