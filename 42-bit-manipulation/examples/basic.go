// Package examples demonstrates Bit Manipulation fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Bit Manipulation example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Bit Manipulation — Basic Example ===\n")
	result := demonstrateBitManipulation()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateBitManipulation() string {
	return "Bit Manipulation basic demonstration"
}
