package arrayhelper

func Paging[T any](arr []T, page, pageSize int) []T {
	if page <= 0 {
		page = 1
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	// Ensure the end index does not exceed the length of the array
	if end > len(arr) {
		end = len(arr)
	}

	// Ensure the start index is not out of bounds
	if start > len(arr) {
		return []T{}
	}

	return arr[start:end]
}
