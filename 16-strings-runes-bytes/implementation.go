// Package strings_runes_bytes implements Strings, Runes & Bytes from scratch for learning purposes.
package strings_runes_bytes

import "errors"

var ErrEmpty = errors.New("strings_runes_bytes: empty structure")

// DataStructure is a placeholder — see module-specific implementations.
type DataStructure struct {
	data []int
}

// New creates a new Strings, Runes & Bytes.
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
