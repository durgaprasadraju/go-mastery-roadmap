// Package examples demonstrates Arrays fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Arrays example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Arrays — Basic Example ===\n")
	result := demonstrateArrays()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateArrays() string {
	return "Arrays basic demonstration"
}
