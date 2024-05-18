package arrayhelper

func At[T any](arr []T, index int) (result T) {
	if index < 0 || index >= len(arr) {
		return result
	}
	return arr[index]
}
