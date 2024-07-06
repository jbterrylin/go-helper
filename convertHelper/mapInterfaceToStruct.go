package converthelper

import "reflect"

func MapInterfaceToStruct[T any](data map[string]interface{}) (result T) {
	t := reflect.TypeOf(result)
	v := reflect.New(t).Elem()

	if v.Kind() != reflect.Struct {
		return
	}

	for key, value := range data {
		field := v.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			val := reflect.ValueOf(value)
			if val.Type().AssignableTo(field.Type()) {
				field.Set(val)
			}
		}
	}

	return v.Interface().(T)
}
