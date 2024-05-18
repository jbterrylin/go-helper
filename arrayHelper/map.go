package arrayhelper

// Map function with named return value
func Map[T any, U any](arr []T, fn func(T) U) (result []U) {
	result = make([]U, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return
}
