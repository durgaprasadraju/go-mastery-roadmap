package solutions

import "errors"

var ErrQueueEmpty = errors.New("queue empty")

// Queue is FIFO.
type Queue struct{ items []int }

func NewQueue() *Queue { return &Queue{} }

func (q *Queue) Enqueue(v int) { q.items = append(q.items, v) }

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, ErrQueueEmpty
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, nil
}