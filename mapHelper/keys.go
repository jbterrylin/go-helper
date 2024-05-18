package maphelper

func Keys[K comparable, V any](m map[K]V) []K {
	var result []K
	for key := range m {
		result = append(result, key)
	}
	return result
}
