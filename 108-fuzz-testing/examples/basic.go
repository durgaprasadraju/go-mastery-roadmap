// Package examples demonstrates Fuzz Testing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Fuzz Testing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Fuzz Testing — Basic Example ===\n")
	result := demonstrateFuzzTesting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFuzzTesting() string {
	return "Fuzz Testing basic demonstration"
}
