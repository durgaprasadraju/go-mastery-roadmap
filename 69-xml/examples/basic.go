// Package examples demonstrates XML fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal XML example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== XML — Basic Example ===\n")
	result := demonstrateXml()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateXml() string {
	return "XML basic demonstration"
}
