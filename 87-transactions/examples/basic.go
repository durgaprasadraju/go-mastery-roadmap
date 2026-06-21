// Package examples demonstrates Transactions fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Transactions example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Transactions — Basic Example ===\n")
	result := demonstrateTransactions()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTransactions() string {
	return "Transactions basic demonstration"
}
