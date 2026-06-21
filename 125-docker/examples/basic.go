// Package examples demonstrates Docker fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Docker example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Docker — Basic Example ===\n")
	result := demonstrateDocker()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDocker() string {
	return "Docker basic demonstration"
}
