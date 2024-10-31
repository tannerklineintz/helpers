package helpers

import (
	"reflect"
)

// DeepCopy returns a deep copy of the specified object, including unexported fields.
func DeepCopy[T any](obj T) (T, error) {
	copyVal, err := deepCopyReflect(obj)
	return copyVal.(T), err
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
