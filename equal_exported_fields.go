package helpers

import "reflect"

// EqualExportedFields compares two structs by only considering exported fields.
func EqualExportedFields(a, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// Ensure both variables are of the same type
	if va.Type() != vb.Type() {
		return false
	}

	// If pointers, get the underlying elements
	if va.Kind() == reflect.Ptr {
		va = va.Elem()
		vb = vb.Elem()
	}

	// Ensure we're comparing structs
	if va.Kind() != reflect.Struct || vb.Kind() != reflect.Struct {
		return false
	}

	t := va.Type()
	for i := 0; i < va.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		fa := va.Field(i)
		fb := vb.Field(i)

		// Recursively compare exported fields
		if !reflect.DeepEqual(fa.Interface(), fb.Interface()) {
			return false
		}
	}
	return true
}
