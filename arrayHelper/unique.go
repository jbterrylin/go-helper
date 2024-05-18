package arrayhelper

func Unique[T comparable](arr []T) []T {
	uniqueMap := make(map[T]bool)
	var result []T
	for _, elem := range arr {
		if _, exists := uniqueMap[elem]; !exists {
			uniqueMap[elem] = true
			result = append(result, elem)
		}
	}
	return result
}
