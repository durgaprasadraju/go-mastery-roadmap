// Package examples demonstrates Divide & Conquer fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Divide & Conquer example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Divide & Conquer — Basic Example ===\n")
	result := demonstrateDivideConquer()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDivideConquer() string {
	return "Divide & Conquer basic demonstration"
}
