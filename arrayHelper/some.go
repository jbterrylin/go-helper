package arrayhelper

func Some[T any](slice []T, callbackFn func(T) bool) bool {
	for _, v := range slice {
		if callbackFn(v) {
			return true
		}
	}
	return false
}
