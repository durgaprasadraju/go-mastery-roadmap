// Package examples demonstrates Maps fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Maps example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Maps — Basic Example ===\n")
	result := demonstrateMaps()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMaps() string {
	return "Maps basic demonstration"
}
