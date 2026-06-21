// Package examples demonstrates GORM ORM fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal GORM ORM example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== GORM ORM — Basic Example ===\n")
	result := demonstrateOrmGorm()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateOrmGorm() string {
	return "GORM ORM basic demonstration"
}
