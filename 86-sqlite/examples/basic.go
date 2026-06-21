// Package examples demonstrates SQLite fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal SQLite example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== SQLite — Basic Example ===\n")
	result := demonstrateSqlite()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSqlite() string {
	return "SQLite basic demonstration"
}
