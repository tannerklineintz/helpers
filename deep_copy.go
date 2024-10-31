package helpers

// import (
// 	"bytes"
// 	"encoding/gob"
// )

import (
	"reflect"
	"unsafe"
)

// DeepCopy returns a deep copy of the specified object, including unexported fields.
func DeepCopy[T any](obj T) T {
	val := reflect.ValueOf(obj)
	valType := val.Type()
	var result reflect.Value

	if val.Kind() == reflect.Ptr {
		result = deepCopyValue(val)
	} else {
		// Create an addressable value for non-pointer types
		addr := reflect.New(valType)
		addr.Elem().Set(val)
		result = deepCopyValue(addr)
		result = result.Elem() // Dereference the pointer to get the value
	}

	return result.Interface().(T)
}

func deepCopyValue(src reflect.Value) reflect.Value {
	srcType := src.Type()

	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return reflect.Zero(srcType)
		}
		dstPtr := reflect.New(srcType.Elem())
		dstPtr.Elem().Set(deepCopyValue(src.Elem()))
		return dstPtr

	case reflect.Interface:
		if src.IsNil() {
			return reflect.Zero(srcType)
		}
		srcElem := src.Elem()
		dstElem := deepCopyValue(srcElem)
		return dstElem.Convert(srcType)

	case reflect.Struct:
		dst := reflect.New(srcType).Elem()
		for i := 0; i < src.NumField(); i++ {
			srcField := src.Field(i)
			dstField := dst.Field(i)

			// Handle unexported fields
			if !srcField.CanSet() {
				if srcField.CanAddr() {
					srcField = reflect.NewAt(srcField.Type(), unsafe.Pointer(srcField.UnsafeAddr())).Elem()
				} else {
					continue // Cannot address the field, skip it
				}
			}

			if !dstField.CanSet() {
				if dstField.CanAddr() {
					dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.UnsafeAddr())).Elem()
				} else {
					continue // Cannot address the field, skip it
				}
			}

			dstField.Set(deepCopyValue(srcField))
		}
		return dst

	case reflect.Slice:
		if src.IsNil() {
			return reflect.Zero(srcType)
		}
		dst := reflect.MakeSlice(srcType, src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			dst.Index(i).Set(deepCopyValue(src.Index(i)))
		}
		return dst

	case reflect.Map:
		if src.IsNil() {
			return reflect.Zero(srcType)
		}
		dst := reflect.MakeMapWithSize(srcType, src.Len())
		for _, key := range src.MapKeys() {
			dst.SetMapIndex(deepCopyValue(key), deepCopyValue(src.MapIndex(key)))
		}
		return dst

	case reflect.Array:
		dst := reflect.New(srcType).Elem()
		for i := 0; i < src.Len(); i++ {
			dst.Index(i).Set(deepCopyValue(src.Index(i)))
		}
		return dst

	default:
		return src
	}
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
