// Package examples demonstrates Redis fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Redis example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Redis — Basic Example ===\n")
	result := demonstrateRedis()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRedis() string {
	return "Redis basic demonstration"
}
