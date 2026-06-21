// Package examples demonstrates Rate Limiting fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Rate Limiting example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Rate Limiting — Basic Example ===\n")
	result := demonstrateRateLimiting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRateLimiting() string {
	return "Rate Limiting basic demonstration"
}
