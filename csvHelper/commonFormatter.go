package csvhelper

import (
	"reflect"
	"strconv"
	"time"
)

func GetDefaultConverterMap(exportTimeZone int64) map[string]*CsvConverter {
	return map[string]*CsvConverter{
		"l_csv__float_to_2_decimal": {
			ConvertFunc: ConvertFloatTo2Decimal,
		},
		"l_csv__convert_time_to_string": {
			ConvertFunc: ConvertTimeToString,
			Value:       exportTimeZone,
		},
	}
}

func ConvertFloatTo2Decimal(data reflect.Value, fieldName string, value reflect.Value, others interface{}) reflect.Value {
	if !value.CanFloat() {
		return value
	}
	tmp := strconv.FormatFloat(value.Float(), 'f', 2, 64)
	val := reflect.ValueOf(tmp)
	return val
}

func ConvertTimeToString(data reflect.Value, fieldName string, value reflect.Value, others interface{}) reflect.Value {
	realValue, ok := value.Interface().(time.Time)
	if ok {
		if others != nil && !realValue.IsZero() {
			offset, ok := others.(int64)
			if ok {
				offsetDuration := time.Duration(offset) * time.Minute
				realValue = realValue.Add(offsetDuration)
				realValueStr := realValue.Format("2006-01-02 15:04:05")
				val := reflect.ValueOf(realValueStr)
				return val
			} else {
				return value
			}
		} else {
			realValueStr := realValue.Format("2006-01-02 15:04:05")
			val := reflect.ValueOf(realValueStr)
			return val
		}
	} else {
		return value
	}
}
