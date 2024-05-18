package mathhelper

func Min[T Comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}
