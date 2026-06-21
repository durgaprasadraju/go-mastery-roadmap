// Package examples demonstrates Channels fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Channels example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Channels — Basic Example ===\n")
	result := demonstrateChannels()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateChannels() string {
	return "Channels basic demonstration"
}
