// Package examples demonstrates Tries fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Tries example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Tries — Basic Example ===\n")
	result := demonstrateTrie()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTrie() string {
	return "Tries basic demonstration"
}
