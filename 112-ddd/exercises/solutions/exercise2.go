package solutions

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
)

// Service demonstrates production patterns for Domain-Driven Design.
type Service struct {
	Name string
}

// Process handles Domain-Driven Design operations with context cancellation.
func (s *Service) Process(ctx context.Context, input string) (string, error) {
	select {
	case <-ctx.Done():
		return "", apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "Domain-Driven Design cancelled")
	default:
	}
	if input == "" {
		return "", apperrors.New(apperrors.CodeValidation, "input required")
	}
	return fmt.Sprintf("[%s] %s: %s", s.Name, "Domain-Driven Design", input), nil
}
