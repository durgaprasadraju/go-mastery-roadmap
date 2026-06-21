// Package examples demonstrates MySQL fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal MySQL example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== MySQL — Basic Example ===\n")
	result := demonstrateMysql()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMysql() string {
	return "MySQL basic demonstration"
}
