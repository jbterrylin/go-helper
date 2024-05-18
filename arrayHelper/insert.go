package arrayhelper

// will panic
func Insert[T any](arr []T, index int, value T) []T {
	if index < 0 || index > len(arr) {
		panic("index out of range")
	}
	arr = append(arr, value) // Make room for the new element
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr
}

// InsertSafe function inserts an element at the given index or at the nearest valid position
func InsertSafe[T any](arr []T, index int, value T) []T {
	if index < 0 {
		index = 0
	} else if index > len(arr) {
		index = len(arr)
	}
	arr = append(arr, value) // Make room for the new element
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr
}

func InsertWithErr[T any](arr []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(arr) {
		return arr, ErrOutOfRange
	}
	arr = append(arr, value) // Make room for the new element
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr, nil
}
