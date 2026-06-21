// Package examples demonstrates SQL Fundamentals fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal SQL Fundamentals example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== SQL Fundamentals — Basic Example ===\n")
	result := demonstrateDatabaseSql()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateDatabaseSql() string {
	return "SQL Fundamentals basic demonstration"
}
