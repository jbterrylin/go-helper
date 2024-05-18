package ophelper

import "reflect"

func Or[T any](value, elseValue T) T {
	// 使用反射检查第一个值是否为零值
	if reflect.DeepEqual(value, *new(T)) {
		return elseValue
	}
	return value
}

func OrByCond[T any](value, elseValue T, condition func(T) bool) T {
	// 如果条件回调函数返回 true 或者第一个值为零值，则返回第二个值
	if condition(value) {
		return elseValue
	}
	return value
}
