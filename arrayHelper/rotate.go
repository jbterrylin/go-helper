package arrayhelper

func Rotate[T any](arr []T, times int) []T {
	n := len(arr)
	if n == 0 {
		return arr
	}

	// 处理旋转次数
	times = times % n
	if times < 0 {
		times += n
	}

	// 旋转数组
	return append(arr[n-times:], arr[:n-times]...)
}

func RotateByIndex[T any](arr []T, index int) []T {
	if index < 0 {
		index = len(arr) + index
	}
	if index < 0 || index >= len(arr) {
		return arr // 如果索引无效，返回原数组
	}

	// 分割数组并重新组合
	rotated := append(arr[index:], arr[:index]...)
	return rotated
}

func RotateByIndexSafe[T any](arr []T, index int) []T {
	if len(arr) == 0 {
		return arr
	}
	if index < 0 {
		index = len(arr) + index
	}
	if index < 0 {
		index = 0
	} else if index >= len(arr) {
		index = len(arr) - 1
	}

	// 分割数组并重新组合
	rotated := append(arr[index:], arr[:index]...)
	return rotated
}
