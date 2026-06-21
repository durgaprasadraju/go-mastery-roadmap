// Package examples demonstrates Hashing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Hashing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Hashing — Basic Example ===\n")
	result := demonstrateHashing()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHashing() string {
	return "Hashing basic demonstration"
}
