package mapcache

import (
	"reflect"
)

// calculateSize estimates the size of an entry value in bytes
func calculateSize(value interface{}) int64 {
	v := reflect.ValueOf(value)
	return calculateSizeRecursive(v)
}

// calculateSizeRecursive calculates the size of the value recursively.
func calculateSizeRecursive(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.Bool:
		return 1
	case reflect.Uint8, reflect.Int8:
		return 1
	case reflect.Int16, reflect.Uint16:
		return 2
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		return 4
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Int, reflect.Uint:
		return 8
	case reflect.String:
		return int64(len(v.String()))
	case reflect.Array, reflect.Slice:
		var size int64 = 0
		for i := 0; i < v.Len(); i++ {
			size += calculateSizeRecursive(v.Index(i))
		}
		return size
	case reflect.Map:
		var size int64 = 0
		for _, key := range v.MapKeys() {
			size += calculateSizeRecursive(key)
			size += calculateSizeRecursive(v.MapIndex(key))
		}
		return size
	case reflect.Struct:
		var size int64 = 0
		for i := 0; i < v.NumField(); i++ {
			size += calculateSizeRecursive(v.Field(i))
		}
		return size
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return 0
		}
		return calculateSizeRecursive(v.Elem())
	default:
		return 0
	}
}
