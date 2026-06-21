// Package examples demonstrates Goroutines fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Goroutines example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Goroutines — Basic Example ===\n")
	result := demonstrateGoroutines()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGoroutines() string {
	return "Goroutines basic demonstration"
}
