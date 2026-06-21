// Package examples demonstrates Trees fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Trees example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Trees — Basic Example ===\n")
	result := demonstrateTrees()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTrees() string {
	return "Trees basic demonstration"
}
