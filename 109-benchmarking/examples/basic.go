// Package examples demonstrates Benchmarking fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Benchmarking example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Benchmarking — Basic Example ===\n")
	result := demonstrateBenchmarking()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateBenchmarking() string {
	return "Benchmarking basic demonstration"
}
