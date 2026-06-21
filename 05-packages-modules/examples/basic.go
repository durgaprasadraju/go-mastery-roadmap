// Package examples demonstrates Packages & Modules fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Packages & Modules example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Packages & Modules — Basic Example ===\n")
	result := demonstratePackagesModules()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstratePackagesModules() string {
	return "Packages & Modules basic demonstration"
}
