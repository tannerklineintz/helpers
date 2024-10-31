package helpers

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

// DeepCopy returns a deep copy of specified object
func DeepCopy(obj any) any {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(obj); err != nil {
		return nil
	}

	// Create a new zero value of the same type
	objType := reflect.TypeOf(obj)
	copyValue := reflect.New(objType).Interface()

	decoder := gob.NewDecoder(&buf)
	if err := decoder.Decode(copyValue); err != nil {
		return nil
	}

	// Dereference pointer if original object is not a pointer
	if objType.Kind() != reflect.Ptr {
		return reflect.ValueOf(copyValue).Elem().Interface()
	}
	return copyValue
}
