// Package examples demonstrates RabbitMQ fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal RabbitMQ example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== RabbitMQ — Basic Example ===\n")
	result := demonstrateRabbitmq()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRabbitmq() string {
	return "RabbitMQ basic demonstration"
}
