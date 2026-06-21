// Package examples demonstrates Networking Fundamentals fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Networking Fundamentals example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Networking Fundamentals — Basic Example ===\n")
	result := demonstrateNetworkingFundamentals()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateNetworkingFundamentals() string {
	return "Networking Fundamentals basic demonstration"
}
