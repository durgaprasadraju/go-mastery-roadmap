// Package examples demonstrates Interview Preparation fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Interview Preparation example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Interview Preparation — Basic Example ===\n")
	result := demonstrateInterviewPreparation()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateInterviewPreparation() string {
	return "Interview Preparation basic demonstration"
}
