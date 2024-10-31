package helpers

import (
	"bytes"
	"encoding/gob"
	"log"
)

// DeepCopy returns a deep copy of specified object
func DeepCopy(obj any) any {
	var buf bytes.Buffer
	var copy any
	encoder, decoder := gob.NewEncoder(&buf), gob.NewDecoder(&buf)
	if err := encoder.Encode(obj); err != nil {
		log.Printf("%v", err.Error())
		return nil
	}
	if err := decoder.Decode(&copy); err != nil {
		log.Printf("%v", err.Error())
		return nil
	}

	return copy
}
