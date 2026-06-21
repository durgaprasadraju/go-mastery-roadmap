// Package examples demonstrates Mocking fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Mocking example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Mocking — Basic Example ===\n")
	result := demonstrateMocking()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMocking() string {
	return "Mocking basic demonstration"
}
