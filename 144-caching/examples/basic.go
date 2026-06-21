// Package examples demonstrates Caching Strategies fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Caching Strategies example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Caching Strategies — Basic Example ===\n")
	result := demonstrateCaching()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCaching() string {
	return "Caching Strategies basic demonstration"
}
