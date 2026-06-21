// Package examples demonstrates Observability fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Observability example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Observability — Basic Example ===\n")
	result := demonstrateObservability()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateObservability() string {
	return "Observability basic demonstration"
}
