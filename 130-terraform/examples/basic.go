// Package examples demonstrates Terraform fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Terraform example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Terraform — Basic Example ===\n")
	result := demonstrateTerraform()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTerraform() string {
	return "Terraform basic demonstration"
}
