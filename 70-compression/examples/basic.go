// Package examples demonstrates Compression fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Compression example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Compression — Basic Example ===\n")
	result := demonstrateCompression()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCompression() string {
	return "Compression basic demonstration"
}
