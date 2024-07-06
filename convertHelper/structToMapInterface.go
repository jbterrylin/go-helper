package converthelper

import "reflect"

// return nil if data is not struct
func StructToMapInterface(data interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})

	// 获取传入数据的值
	v := reflect.ValueOf(data)

	// 确保传入的是一个结构体
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	// 获取结构体的类型
	t := v.Type()

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		result[t.Field(i).Name] = v.Field(i).Interface()
	}

	return
}
