// Package examples demonstrates Distributed Tracing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Distributed Tracing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Distributed Tracing — Basic Example ===\n")
	result := demonstrateTracing()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTracing() string {
	return "Distributed Tracing basic demonstration"
}
