package solutions

// SlidingWindow limits requests per window size.
type SlidingWindow struct {
	limit int
	count int
}

func NewSlidingWindow(limit int) *SlidingWindow {
	return &SlidingWindow{limit: limit}
}

func (sw *SlidingWindow) Allow() bool {
	if sw.count >= sw.limit {
		return false
	}
	sw.count++
	return true
}