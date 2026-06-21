// Package examples demonstrates Dynamic Programming fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Dynamic Programming example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Dynamic Programming — Basic Example ===\n")
	result := demonstrateDynamicProgramming()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDynamicProgramming() string {
	return "Dynamic Programming basic demonstration"
}
