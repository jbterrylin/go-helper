package arrayhelper

import (
	mathhelper "github.com/jbterrylin/go-helper/mathHelper"
)

func Fill[T any](arr []T, value T, startEnd ...int) []T {
	start := 0
	end := len(arr)

	// 处理可选参数
	switch len(startEnd) {
	case 1:
		start = startEnd[0]
	case 2:
		start = startEnd[0]
		end = startEnd[1]
	}

	// 规范化索引
	if start < 0 {
		start = mathhelper.Max(len(arr)+start, 0)
	} else {
		start = mathhelper.Min(start, len(arr))
	}

	if end < 0 {
		end = mathhelper.Max(len(arr)+end, 0)
	} else {
		end = mathhelper.Min(end, len(arr))
	}

	// 填充元素
	for i := start; i < end; i++ {
		arr[i] = value
	}

	return arr
}

// FillWithExtensible 填充数组中的元素并在需要时扩展数组长度
func FillWithExtensible[T any](arr []T, value T, startEnd ...int) []T {
	start := 0
	end := len(arr)

	// 处理可选参数
	switch len(startEnd) {
	case 1:
		start = startEnd[0]
	case 2:
		start = startEnd[0]
		end = startEnd[1]
	}

	// 规范化索引
	if start < 0 {
		start = mathhelper.Max(len(arr)+start, 0)
	}
	if end < 0 {
		end = mathhelper.Max(len(arr)+end, 0)
	}

	// 如果 end 超出数组长度，扩展数组
	if end > len(arr) {
		extended := make([]T, end)
		copy(extended, arr)
		arr = extended
		end = len(arr)
	}

	// Ensure start index is within bounds
	if start > len(arr) {
		extended := make([]T, start+1)
		copy(extended, arr)
		arr = extended
		end = len(arr)
	}

	// Fill elements
	for i := start; i < end; i++ {
		arr[i] = value
	}

	return arr
}
