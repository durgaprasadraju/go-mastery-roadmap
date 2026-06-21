// Package examples demonstrates RWMutex fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal RWMutex example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== RWMutex — Basic Example ===\n")
	result := demonstrateRwmutex()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRwmutex() string {
	return "RWMutex basic demonstration"
}
