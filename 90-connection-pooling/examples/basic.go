// Package examples demonstrates Connection Pooling fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Connection Pooling example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Connection Pooling — Basic Example ===\n")
	result := demonstrateConnectionPooling()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateConnectionPooling() string {
	return "Connection Pooling basic demonstration"
}
