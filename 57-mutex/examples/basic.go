// Package examples demonstrates Mutex fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Mutex example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Mutex — Basic Example ===\n")
	result := demonstrateMutex()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMutex() string {
	return "Mutex basic demonstration"
}
