package arrayhelper

func Shift[T any](slice *[]T) T {
	if len(*slice) == 0 {
		panic(ErrOutOfRange.Error())
	}
	elem := (*slice)[0]
	*slice = (*slice)[:0]
	return elem
}

func ShiftWithErr[T any](slice *[]T) (T, error) {
	if len(*slice) == 0 {
		return *new(T), ErrOutOfRange
	}
	elem := (*slice)[0]
	*slice = (*slice)[:0]
	return elem, nil
}
