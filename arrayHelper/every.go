package arrayhelper

func Every[T any](arr []T, test func(T) bool) bool {
	for _, v := range arr {
		if !test(v) {
			return false
		}
	}
	return true
}
