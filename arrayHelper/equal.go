package arrayhelper

func Equal[T comparable](firstArr, secondArr []T) bool {
	if len(firstArr) != len(secondArr) {
		return false
	}

	for i := range firstArr {
		if firstArr[i] != secondArr[i] {
			return false
		}
	}
	return true
}

func EqualRef[T comparable](firstArr, secondArr []*T) bool {
	if len(firstArr) != len(secondArr) {
		return false
	}

	for i := range firstArr {
		if firstArr[i] == nil || secondArr[i] == nil {
			if firstArr[i] != secondArr[i] {
				return false
			}
		} else if *firstArr[i] != *secondArr[i] {
			return false
		}
	}
	return true
}
