package solutions

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
)

// Service demonstrates production patterns for Two Pointers.
type Service struct {
	Name string
}

// Process handles Two Pointers operations with context cancellation.
func (s *Service) Process(ctx context.Context, input string) (string, error) {
	select {
	case <-ctx.Done():
		return "", apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "Two Pointers cancelled")
	default:
	}
	if input == "" {
		return "", apperrors.New(apperrors.CodeValidation, "input required")
	}
	return fmt.Sprintf("[%s] %s: %s", s.Name, "Two Pointers", input), nil
}
