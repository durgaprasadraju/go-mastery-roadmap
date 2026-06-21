// Package examples demonstrates Load Balancing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Load Balancing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Load Balancing — Basic Example ===\n")
	result := demonstrateLoadBalancing()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateLoadBalancing() string {
	return "Load Balancing basic demonstration"
}
