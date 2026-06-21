// Package examples demonstrates HTTP fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal HTTP example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== HTTP — Basic Example ===\n")
	result := demonstrateHttp()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHttp() string {
	return "HTTP basic demonstration"
}
