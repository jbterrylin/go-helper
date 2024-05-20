package validhelper

import (
	"reflect"
	"time"
)

func isTimeType(val reflect.Value) bool {
	// 检查是否是 time.Time 类型
	return val.Type() == reflect.TypeOf(time.Time{})
}

func timeToUnixTime(val reflect.Value) int64 {
	// func toUnixTime(val reflect.Value) (int64, error) {
	// if !isTimeType(val) {
	// 	return 0, fmt.Errorf("value is not of type time.Time")
	// }
	// 将 reflect.Value 转换回 time.Time
	t := val.Interface().(time.Time)
	// 获取 Unix 时间
	return t.Unix()
}
