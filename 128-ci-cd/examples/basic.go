// Package examples demonstrates CI/CD fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal CI/CD example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== CI/CD — Basic Example ===\n")
	result := demonstrateCiCd()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCiCd() string {
	return "CI/CD basic demonstration"
}
