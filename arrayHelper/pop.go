package arrayhelper

func Pop[T any](slice *[]T) T {
	if len(*slice) == 0 {
		panic(ErrOutOfRange.Error())
	}
	elem := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return elem
}

func PopWithErr[T any](slice *[]T) (T, error) {
	if len(*slice) == 0 {
		return *new(T), ErrOutOfRange
	}
	elem := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return elem, nil
}
