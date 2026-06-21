// Package examples demonstrates OpenTelemetry fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal OpenTelemetry example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== OpenTelemetry — Basic Example ===\n")
	result := demonstrateOpentelemetry()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateOpentelemetry() string {
	return "OpenTelemetry basic demonstration"
}
