// Package examples demonstrates Domain-Driven Design fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Domain-Driven Design example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Domain-Driven Design — Basic Example ===\n")
	result := demonstrateDdd()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDdd() string {
	return "Domain-Driven Design basic demonstration"
}
