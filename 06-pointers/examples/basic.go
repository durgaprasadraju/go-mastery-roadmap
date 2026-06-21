// Package examples demonstrates Pointers fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Pointers example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Pointers — Basic Example ===\n")
	result := demonstratePointers()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstratePointers() string {
	return "Pointers basic demonstration"
}
