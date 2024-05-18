package arrayhelper

func FindIndex[T any](arr []T, test func(T) bool) int {
	for i, v := range arr {
		if test(v) {
			return i
		}
	}
	return -1
}

func FindNthIndex[T any](arr []T, test func(T) bool, findNth int, force bool) int {
	var foundIndexes []int
	for i, v := range arr {
		if test(v) {
			foundIndexes = append(foundIndexes, i)
		}
		if len(foundIndexes) == findNth && !force {
			return foundIndexes[len(foundIndexes)-1]
		}
	}

	// 处理负索引
	if findNth < 0 {
		positiveIndex := len(foundIndexes) + findNth
		if positiveIndex >= 0 && positiveIndex < len(foundIndexes) {
			return foundIndexes[positiveIndex]
		}
		return -1
	}

	// 在force为true时，如果找不到足够的匹配项
	if force && len(foundIndexes) < findNth {
		return -1
	}

	// 返回第findNth个匹配项的索引
	if len(foundIndexes) >= findNth {
		return foundIndexes[findNth-1]
	}

	// 如果没有找到足够的匹配项，返回-1
	return -1
}
