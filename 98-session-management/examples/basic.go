// Package examples demonstrates Session Management fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Session Management example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Session Management — Basic Example ===\n")
	result := demonstrateSessionManagement()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSessionManagement() string {
	return "Session Management basic demonstration"
}
