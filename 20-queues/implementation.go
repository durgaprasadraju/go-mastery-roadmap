// Package queues implements queue and deque data structures from scratch.
package queues

import "errors"

var ErrEmpty = errors.New("queue is empty")

// Queue is a FIFO queue using a circular buffer approach via slice.
type Queue struct {
	items []int
}

// New creates an empty queue.
func New() *Queue {
	return &Queue{items: make([]int, 0)}
}

// Enqueue adds to the rear. O(1) amortized.
func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

// Dequeue removes from the front. O(n) due to slice shift — production uses ring buffer.
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, ErrEmpty
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, nil
}

// Front returns the front element. O(1).
func (q *Queue) Front() (int, error) {
	if len(q.items) == 0 {
		return 0, ErrEmpty
	}
	return q.items[0], nil
}

// Size returns element count.
func (q *Queue) Size() int { return len(q.items) }

// Deque is a double-ended queue.
type Deque struct {
	items []int
}

// NewDeque creates an empty deque.
func NewDeque() *Deque {
	return &Deque{items: make([]int, 0)}
}

// PushFront inserts at front.
func (d *Deque) PushFront(v int) {
	d.items = append([]int{v}, d.items...)
}

// PushBack inserts at back.
func (d *Deque) PushBack(v int) {
	d.items = append(d.items, v)
}

// PopFront removes from front.
func (d *Deque) PopFront() (int, error) {
	if len(d.items) == 0 {
		return 0, ErrEmpty
	}
	v := d.items[0]
	d.items = d.items[1:]
	return v, nil
}

// PopBack removes from back.
func (d *Deque) PopBack() (int, error) {
	if len(d.items) == 0 {
		return 0, ErrEmpty
	}
	last := len(d.items) - 1
	v := d.items[last]
	d.items = d.items[:last]
	return v, nil
}
