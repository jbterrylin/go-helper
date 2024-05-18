package arrayhelper

func Find[T any](arr []T, test func(T) bool) T {
	for _, v := range arr {
		if test(v) {
			return v
		}
	}
	// 返回类型 T 的零值
	return *new(T)
}

func FindNth[T any](arr []T, test func(T) bool, findNth int) T {
	var found []T
	for _, v := range arr {
		if test(v) {
			found = append(found, v)
		}
		if len(found) == findNth {
			return found[len(found)-1]
		}
	}
	if len(found) < findNth {
		// 返回类型 T 的零值
		return *new(T)
	}
	// 处理负索引
	if findNth < 0 {
		positiveIndex := len(found) + findNth
		if positiveIndex < 0 || positiveIndex >= len(found) {
			return *new(T)
		}
		return found[positiveIndex]
	}
	if len(found) >= findNth {
		return found[findNth-1]
	}
	// 返回类型 T 的零值
	return *new(T)
}

func FindWithErr[T any](arr []T, test func(T) bool) (T, error) {
	for _, v := range arr {
		if test(v) {
			return v, nil
		}
	}
	return *new(T), ErrNotFound
}

func FindNthWithErr[T any](arr []T, test func(T) bool, findNth int) (T, error) {
	var found []T
	for _, v := range arr {
		if test(v) {
			found = append(found, v)
		}
		if len(found) == findNth {
			return found[len(found)-1], nil
		}
	}
	if len(found) < findNth {
		return *new(T), ErrNotFound
	}
	// 处理负索引
	if findNth < 0 {
		positiveIndex := len(found) + findNth
		if positiveIndex < 0 || positiveIndex >= len(found) {
			return *new(T), ErrNotFound
		}
		return found[positiveIndex], nil
	}
	if len(found) >= findNth {
		return found[findNth-1], nil
	}
	return *new(T), ErrNotFound
}
