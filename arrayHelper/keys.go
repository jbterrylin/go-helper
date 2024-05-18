package arrayhelper

func Keys[T any](arr []T) []int {
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = i
	}
	return result
}
