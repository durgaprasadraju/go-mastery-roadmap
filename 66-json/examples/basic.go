// Package examples demonstrates JSON fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal JSON example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== JSON — Basic Example ===\n")
	result := demonstrateJson()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateJson() string {
	return "JSON basic demonstration"
}
