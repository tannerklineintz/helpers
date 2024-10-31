package helpers

import (
	"bytes"
	"encoding/gob"
)

// DeepCopy returns a deep copy of specified object
func DeepCopy(obj any) any {
	var buf bytes.Buffer
	var copy any
	encoder, decoder := gob.NewEncoder(&buf), gob.NewDecoder(&buf)
	if err := encoder.Encode(obj); err != nil {
		return nil
	}
	if err := decoder.Decode(copy); err != nil {
		return nil
	}

	return copy
}
