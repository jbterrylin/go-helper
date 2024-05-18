package arrayhelper

func GroupBy[T any, U comparable](arr []T, getKey func(T) U) (result map[U][]T) {
	result = map[U][]T{}
	for _, v := range arr {
		key := getKey(v)
		_, exist := result[key]
		if exist {
			result[key] = append(result[key], v)
		} else {
			result[key] = []T{v}
		}
	}
	return
}

func GroupByAndReshape[T any, U comparable, V any](arr []T, getKey func(T) U, reshape func(T) V) (result map[U][]V) {
	result = map[U][]V{}
	for _, v := range arr {
		key := getKey(v)
		_, exist := result[key]
		if exist {
			result[key] = append(result[key], reshape(v))
		} else {
			result[key] = []V{reshape(v)}
		}
	}
	return
}
