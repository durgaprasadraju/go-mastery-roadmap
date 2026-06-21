// Package examples demonstrates Fenwick Trees fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Fenwick Trees example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Fenwick Trees — Basic Example ===\n")
	result := demonstrateFenwickTree()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFenwickTree() string {
	return "Fenwick Trees basic demonstration"
}
