// Package examples demonstrates NATS fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal NATS example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== NATS — Basic Example ===\n")
	result := demonstrateNats()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateNats() string {
	return "NATS basic demonstration"
}
