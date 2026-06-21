// Package examples demonstrates Big-O Analysis fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Big-O Analysis example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Big-O Analysis — Basic Example ===\n")
	result := demonstrateBigOAnalysis()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateBigOAnalysis() string {
	return "Big-O Analysis basic demonstration"
}
