// Package examples demonstrates DNS fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal DNS example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== DNS — Basic Example ===\n")
	result := demonstrateDns()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDns() string {
	return "DNS basic demonstration"
}
