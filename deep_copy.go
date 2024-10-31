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
	copyVal := deepCopyValue(val)
	return copyVal.Interface().(T)
}

func deepCopyValue(src reflect.Value) reflect.Value {
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dstPtr := reflect.New(src.Type().Elem())
		dstPtr.Elem().Set(deepCopyValue(src.Elem()))
		return dstPtr
	case reflect.Interface:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		return deepCopyValue(src.Elem())
	case reflect.Struct:
		dst := reflect.New(src.Type()).Elem()
		for i := 0; i < src.NumField(); i++ {
			srcField := src.Field(i)
			dstField := dst.Field(i)

			// Handle unexported fields
			if !srcField.CanInterface() {
				srcField = reflect.NewAt(srcField.Type(), unsafe.Pointer(srcField.UnsafeAddr())).Elem()
			}
			if !dstField.CanSet() {
				dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.UnsafeAddr())).Elem()
			}

			dstField.Set(deepCopyValue(srcField))
		}
		return dst
	case reflect.Slice:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeSlice(src.Type(), src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			dst.Index(i).Set(deepCopyValue(src.Index(i)))
		}
		return dst
	case reflect.Map:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeMapWithSize(src.Type(), src.Len())
		for _, key := range src.MapKeys() {
			dst.SetMapIndex(deepCopyValue(key), deepCopyValue(src.MapIndex(key)))
		}
		return dst
	case reflect.Array:
		dst := reflect.New(src.Type()).Elem()
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
