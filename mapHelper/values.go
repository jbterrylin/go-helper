package maphelper

func Values[K comparable, V any](m map[K]V) []V {
	var result []V
	for key := range m {
		result = append(result, m[key])
	}
	return result
}
