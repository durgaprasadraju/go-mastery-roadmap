package solutions

import (
	"context"
	"time"
)

// WithTimeoutOp runs until context deadline.
func WithTimeoutOp(ctx context.Context, ms int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(time.Duration(ms*2) * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}