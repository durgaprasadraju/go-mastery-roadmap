package solutions

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
)

// Exercise2Service demonstrates production patterns for Methods.
type Exercise2Service struct {
	Topic string
}

// Process runs with context cancellation support.
func (s *Exercise2Service) Process(ctx context.Context, data string) (string, error) {
	select {
	case <-ctx.Done():
		return "", apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "Methods processing cancelled")
	default:
	}
	if data == "" {
		return "", apperrors.New(apperrors.CodeValidation, "data is required")
	}
	return fmt.Sprintf("[%s] processed: %s", s.Topic, data), nil
}
