package arrayhelper

// Reduce function with named return value
func Reduce[T any, U any](arr []T, fn func(U, T) U, initial U) (result U) {
	result = initial
	for _, v := range arr {
		result = fn(result, v)
	}
	return
}
