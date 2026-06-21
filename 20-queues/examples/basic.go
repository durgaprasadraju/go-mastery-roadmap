// Package examples demonstrates Queues fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Queues example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Queues — Basic Example ===\n")
	result := demonstrateQueues()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateQueues() string {
	return "Queues basic demonstration"
}
