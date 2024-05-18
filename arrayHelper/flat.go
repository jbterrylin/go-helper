package arrayhelper

import "reflect"

func Flatten[T any](input interface{}) []T {
	var result []T
	val := reflect.ValueOf(input)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			innerVal := val.Index(i).Interface()
			result = append(result, Flatten[T](innerVal)...)
		}
	default:
		result = append(result, input.(T))
	}

	return result
}

func FlattenAll[T any](input interface{}) []T {
	var result []T
	val := reflect.ValueOf(input)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			innerVal := val.Index(i).Interface()
			result = append(result, FlattenAll[T](innerVal)...)
		}
	default:
		result = append(result, input.(T))
	}

	return result
}
