// Package priority_queue implements Priority Queues from scratch for learning purposes.
package priority_queue

import "errors"

var ErrEmpty = errors.New("priority_queue: empty structure")

// DataStructure is a placeholder — see module-specific implementations.
type DataStructure struct {
	data []int
}

// New creates a new Priority Queues.
func New() *DataStructure {
	return &DataStructure{data: make([]int, 0)}
}

// Size returns the number of elements.
func (ds *DataStructure) Size() int {
	return len(ds.data)
}

// IsEmpty reports whether the structure has no elements.
func (ds *DataStructure) IsEmpty() bool {
	return len(ds.data) == 0
}
