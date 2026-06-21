// Package examples demonstrates CSV fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal CSV example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== CSV — Basic Example ===\n")
	result := demonstrateCsv()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCsv() string {
	return "CSV basic demonstration"
}
