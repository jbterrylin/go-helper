package arrayhelper

func Splice[T any](slice *[]T, start int, deleteCount int, items ...T) []T {
	// 获取切片的长度
	length := len(*slice)

	// 处理负的 start 参数
	if start < 0 {
		start = length + start
		if start < 0 {
			start = 0
		}
	}

	// 限制 start 的最大值
	if start > length {
		start = length
	}

	// 处理 deleteCount 参数
	if deleteCount < 0 {
		deleteCount = 0
	}

	if start+deleteCount > length {
		deleteCount = length - start
	}

	// 创建 removed 的副本
	removed := append([]T(nil), (*slice)[start:start+deleteCount]...)

	// 分割切片
	remaining := append((*slice)[:start], append(items, (*slice)[start+deleteCount:]...)...)

	// 更新原始切片
	*slice = remaining

	return removed
}
