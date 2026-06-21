// Package examples demonstrates Microservices fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Microservices example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Microservices — Basic Example ===\n")
	result := demonstrateMicroservices()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMicroservices() string {
	return "Microservices basic demonstration"
}
