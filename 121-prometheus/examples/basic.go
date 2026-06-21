// Package examples demonstrates Prometheus fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Prometheus example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Prometheus — Basic Example ===\n")
	result := demonstratePrometheus()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstratePrometheus() string {
	return "Prometheus basic demonstration"
}
