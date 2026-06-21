// Package examples demonstrates Event Streaming fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Event Streaming example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Event Streaming — Basic Example ===\n")
	result := demonstrateEventStreaming()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateEventStreaming() string {
	return "Event Streaming basic demonstration"
}
