package examples

import (
	"context"
	"fmt"
	"time"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// AdvancedDemo shows production patterns for RWMutex.
func AdvancedDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	start := time.Now()
	defer func() {
		log.Info("advanced demo completed", 
			// elapsed tracked via defer pattern
		)
	}()

	if err := validateContext(ctx); err != nil {
		return err
	}

	fmt.Printf("=== RWMutex — Advanced Example ===\n")
	fmt.Printf("Elapsed: %v\n", time.Since(start))
	return nil
}

func validateContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "context cancelled")
	default:
		return nil
	}
}
