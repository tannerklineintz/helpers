package helpers

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

func DeepCopy[T any](obj T) (T, error) {
	var copy T
	// Option 1: Using JSON marshaling (simpler but slower)
	// data, err := json.Marshal(obj)
	// if err != nil {
	// 	return copy, err
	// }
	// err = json.Unmarshal(data, &copy)
	// return copy, err
	// Option 2: Using reflection (more complex but handles more cases)
	copyVal, err := deepCopyReflect(obj)
	copy = copyVal.(T)
	return copy, err
}

// Alternative reflection-based implementation
func deepCopyReflect(obj interface{}) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}
	original := reflect.ValueOf(obj)
	if original.Kind() != reflect.Ptr && original.Kind() != reflect.Interface {
		return obj, nil // Return immutable types as-is
	}
	copy := reflect.New(original.Type().Elem())
	copyRecursive(original.Elem(), copy.Elem())
	return copy.Interface(), nil
}

func copyRecursive(original, copy reflect.Value) {
	switch original.Kind() {
	case reflect.Ptr:
		if !original.IsNil() {
			copy.Set(reflect.New(original.Type().Elem()))
			copyRecursive(original.Elem(), copy.Elem())
		}
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			if original.Field(i).CanInterface() {
				copyRecursive(original.Field(i), copy.Field(i))
			}
		}
	case reflect.Slice:
		if !original.IsNil() {
			copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
			for i := 0; i < original.Len(); i++ {
				copyRecursive(original.Index(i), copy.Index(i))
			}
		}
	case reflect.Map:
		if !original.IsNil() {
			copy.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				copyValue := reflect.New(original.MapIndex(key).Type()).Elem()
				copyRecursive(original.MapIndex(key), copyValue)
				copy.SetMapIndex(key, copyValue)
			}
		}
	default:
		if original.CanInterface() {
			copy.Set(original)
		}
	}
}

// DeepCopy returns a deep copy of the specified object, including unexported fields.
// func DeepCopy[T any](obj T) T {
// 	var copyVal T

// 	// Get the size of the object in bytes
// 	size := unsafe.Sizeof(obj)

// 	// Get pointers to the source and destination objects
// 	srcPtr := unsafe.Pointer(&obj)
// 	dstPtr := unsafe.Pointer(&copyVal)

// 	// Create byte slices pointing to the memory of the objects
// 	srcSlice := unsafe.Slice((*byte)(srcPtr), size)
// 	dstSlice := unsafe.Slice((*byte)(dstPtr), size)

// 	// Copy the memory from the source to the destination
// 	copy(dstSlice, srcSlice)

// 	return copyVal
// }

// DeepCopy returns a deep copy of specified object with only the exported values
func DeepCopyExported[T any](obj T) T {
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
