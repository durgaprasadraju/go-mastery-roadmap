// Package examples demonstrates Monoliths fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Monoliths example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Monoliths — Basic Example ===\n")
	result := demonstrateMonoliths()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMonoliths() string {
	return "Monoliths basic demonstration"
}
