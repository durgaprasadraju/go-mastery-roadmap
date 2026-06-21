// Package fast_slow_pointers implements Fast & Slow Pointers from scratch for learning purposes.
package fast_slow_pointers

import "errors"

var ErrEmpty = errors.New("fast_slow_pointers: empty structure")

// DataStructure is a placeholder — see module-specific implementations.
type DataStructure struct {
	data []int
}

// New creates a new Fast & Slow Pointers.
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
