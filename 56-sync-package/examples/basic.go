// Package examples demonstrates sync Package fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal sync Package example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== sync Package — Basic Example ===\n")
	result := demonstrateSyncPackage()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSyncPackage() string {
	return "sync Package basic demonstration"
}
