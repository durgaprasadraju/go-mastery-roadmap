// Package examples demonstrates Atomic Operations fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Atomic Operations example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Atomic Operations — Basic Example ===\n")
	result := demonstrateAtomic()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateAtomic() string {
	return "Atomic Operations basic demonstration"
}
