package helpers

import (
	"bytes"
	"encoding/gob"
)

// DeepCopy returns a deep copy of specified object
func DeepCopy[T any](obj T) T {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(obj); err != nil {
		var zero T
		return zero
	}

	var copy T
	decoder := gob.NewDecoder(&buf)
	if err := decoder.Decode(&copy); err != nil {
		var zero T
		return zero
	}

	return copy
}
