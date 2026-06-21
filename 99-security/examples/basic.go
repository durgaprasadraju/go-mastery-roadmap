// Package examples demonstrates Security Fundamentals fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Security Fundamentals example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Security Fundamentals — Basic Example ===\n")
	result := demonstrateSecurity()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSecurity() string {
	return "Security Fundamentals basic demonstration"
}
