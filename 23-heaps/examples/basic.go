// Package examples demonstrates Heaps fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Heaps example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Heaps — Basic Example ===\n")
	result := demonstrateHeaps()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHeaps() string {
	return "Heaps basic demonstration"
}
