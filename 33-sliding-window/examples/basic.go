// Package examples demonstrates Sliding Window fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Sliding Window example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Sliding Window — Basic Example ===\n")
	result := demonstrateSlidingWindow()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSlidingWindow() string {
	return "Sliding Window basic demonstration"
}
