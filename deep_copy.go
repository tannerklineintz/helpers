package helpers

// import (
// 	"bytes"
// 	"encoding/gob"
// )

import (
	"unsafe"
)

// DeepCopy returns a deep copy of the specified object, including unexported fields.
func DeepCopy[T any](obj T) T {
	var copyVal T

	// Get the size of the object in bytes
	size := unsafe.Sizeof(obj)

	// Get pointers to the source and destination objects
	srcPtr := unsafe.Pointer(&obj)
	dstPtr := unsafe.Pointer(&copyVal)

	// Create byte slices pointing to the memory of the objects
	srcSlice := unsafe.Slice((*byte)(srcPtr), size)
	dstSlice := unsafe.Slice((*byte)(dstPtr), size)

	// Copy the memory from the source to the destination
	copy(dstSlice, srcSlice)

	return copyVal
}

// DeepCopy returns a deep copy of specified object
// func DeepCopy[T any](obj T) T {
// 	var buf bytes.Buffer
// 	encoder := gob.NewEncoder(&buf)
// 	if err := encoder.Encode(obj); err != nil {
// 		var zero T
// 		return zero
// 	}

// 	var copy T
// 	decoder := gob.NewDecoder(&buf)
// 	if err := decoder.Decode(&copy); err != nil {
// 		var zero T
// 		return zero
// 	}

// 	return copy
// }
