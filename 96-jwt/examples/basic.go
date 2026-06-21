// Package examples demonstrates JWT fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal JWT example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== JWT — Basic Example ===\n")
	result := demonstrateJwt()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateJwt() string {
	return "JWT basic demonstration"
}
