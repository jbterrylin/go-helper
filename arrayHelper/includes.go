package arrayhelper

func Includes[T comparable](slice []T, searchElement T, fromIndex ...int) bool {
	start := 0
	if len(fromIndex) > 0 {
		start = fromIndex[0]
		if start < 0 {
			start = len(slice) + start
		}
		if start < 0 {
			start = 0
		}
	}
	for i := start; i < len(slice); i++ {
		if slice[i] == searchElement {
			return true
		}
	}
	return false
}
