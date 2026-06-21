// Package examples demonstrates Apache Kafka fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Apache Kafka example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Apache Kafka — Basic Example ===\n")
	result := demonstrateKafka()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateKafka() string {
	return "Apache Kafka basic demonstration"
}
