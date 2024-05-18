package arrayhelper

func Init[T any](length int, setValue func(int) T) []T {
	arr := make([]T, length)
	for i := range arr {
		arr[i] = setValue(i)
	}
	return arr
}

func InitConst[T any](length int, value T) []T {
	arr := make([]T, length)
	for i := range arr {
		arr[i] = value
	}
	return arr
}
