package solutions

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
)

// Service demonstrates production patterns for HTTP/2.
type Service struct {
	Name string
}

// Process handles HTTP/2 operations with context cancellation.
func (s *Service) Process(ctx context.Context, input string) (string, error) {
	select {
	case <-ctx.Done():
		return "", apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "HTTP/2 cancelled")
	default:
	}
	if input == "" {
		return "", apperrors.New(apperrors.CodeValidation, "input required")
	}
	return fmt.Sprintf("[%s] %s: %s", s.Name, "HTTP/2", input), nil
}
