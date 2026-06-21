// Package examples demonstrates Hexagonal Architecture fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Hexagonal Architecture example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Hexagonal Architecture — Basic Example ===\n")
	result := demonstrateHexagonal()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHexagonal() string {
	return "Hexagonal Architecture basic demonstration"
}
