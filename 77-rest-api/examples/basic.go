// Package examples demonstrates REST API fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal REST API example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== REST API — Basic Example ===\n")
	result := demonstrateRestApi()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRestApi() string {
	return "REST API basic demonstration"
}
