package converthelper

import (
	"fmt"
	"reflect"
)

func FlatStructToMapString(data interface{}) (result map[string]string) {
	result = make(map[string]string)

	// 获取传入数据的值
	v := reflect.ValueOf(data)

	// 确保传入的是一个结构体
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}

	// 获取结构体的类型
	t := v.Type()

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// 将字段值转换为字符串并存储到结果 map 中
		result[fieldName] = fmt.Sprintf("%v", field.Interface())
	}

	return result
}

func NestedStructToMapString(data interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})

	v := reflect.ValueOf(data)

	// 确保传入的是一个结构体
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// 将字段值转换为字符串并存储到结果 map 中
		if field.Kind() == reflect.Ptr {
			if !field.IsNil() {
				field = field.Elem()
				if field.Kind() == reflect.Struct {
					result[fieldName] = NestedStructToMapString(field.Interface())
				} else {
					result[fieldName] = fmt.Sprintf("%v", field.Interface())
				}
			}
		} else {
			if field.Kind() == reflect.Struct {
				result[fieldName] = NestedStructToMapString(field.Interface())
			} else {
				result[fieldName] = fmt.Sprintf("%v", field.Interface())
			}
		}
	}

	return result
}
