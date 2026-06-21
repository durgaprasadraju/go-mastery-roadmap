// Package examples demonstrates Event-Driven Architecture fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Event-Driven Architecture example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Event-Driven Architecture — Basic Example ===\n")
	result := demonstrateEventDriven()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateEventDriven() string {
	return "Event-Driven Architecture basic demonstration"
}
