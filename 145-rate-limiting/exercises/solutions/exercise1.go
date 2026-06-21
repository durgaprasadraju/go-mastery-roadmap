package solutions

// TokenBucket rate limiter (simplified).
type TokenBucket struct{ tokens, cap int }

func NewTokenBucket(cap int) *TokenBucket { return &TokenBucket{tokens: cap, cap: cap} }

func (tb *TokenBucket) Allow() bool {
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}