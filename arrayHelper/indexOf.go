package arrayhelper

// indexOf 函数在切片中查找元素的索引，支持从指定索引开始查找
func IndexOf[T comparable](slice []T, searchElement T, fromIndex ...int) int {
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
			return i
		}
	}
	return -1
}
