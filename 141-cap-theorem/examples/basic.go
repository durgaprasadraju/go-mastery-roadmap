// Package examples demonstrates CAP Theorem fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal CAP Theorem example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== CAP Theorem — Basic Example ===\n")
	result := demonstrateCapTheorem()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCapTheorem() string {
	return "CAP Theorem basic demonstration"
}
