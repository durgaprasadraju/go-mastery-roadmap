// Package examples demonstrates TCP fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal TCP example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== TCP — Basic Example ===\n")
	result := demonstrateTcp()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTcp() string {
	return "TCP basic demonstration"
}
