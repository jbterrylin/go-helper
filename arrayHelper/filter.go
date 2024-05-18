package arrayhelper

// Filter function creates a new array with all elements that pass the test implemented by the provided function.
func Filter[T any](arr []T, fn func(T) bool) []T {
	result := []T{}
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
