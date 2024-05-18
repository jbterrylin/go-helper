package arrayhelper

func ValueEqual[T comparable](firstArr, secondArr []T) bool {
	if len(firstArr) != len(secondArr) {
		return false
	}

	// Create a copy of secondArr to modify
	secondArrCopy := make([]T, len(secondArr))
	copy(secondArrCopy, secondArr)

	for _, val := range firstArr {
		index := IndexOf(secondArrCopy, val)
		if index != -1 {
			Splice(&secondArrCopy, index, 1)
		} else {
			return false
		}
	}
	return true
}
