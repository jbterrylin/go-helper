package csvhelper

import (
	"reflect"
	"strconv"
	"strings"
)

func getTagValue(tag, key string) string {
	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, key) {
			return strings.TrimPrefix(part, key)
		}
	}
	return ""
}

func CsvToStructs(data [][]string, values reflect.Value) error {
	if values.Kind() != reflect.Ptr || values.Elem().Kind() != reflect.Slice {
		return ErrValsNotSlicePointer
	}

	headers := data[0]
	slice := values.Elem()
	structType := slice.Type().Elem()

	for i, row := range data {
		if i == 0 { // Skip headers
			continue
		}

		newStruct := reflect.New(structType).Elem()

		for j, cell := range row {
			header := headers[j]

			for k := 0; k < newStruct.NumField(); k++ {
				field := newStruct.Type().Field(k)
				tagHeader := getTagValue(field.Tag.Get(CsvTag), CsvHeader)
				if tagHeader == "" {
					tagHeader = field.Name
				}

				if tagHeader == header {
					switch newStruct.Field(k).Kind() {
					case reflect.String:
						newStruct.Field(k).SetString(cell)
					case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int:
						val, _ := strconv.ParseInt(cell, 10, 64)
						newStruct.Field(k).SetInt(val)
					case reflect.Float64, reflect.Float32:
						val, _ := strconv.ParseFloat(cell, 64)
						newStruct.Field(k).SetFloat(val)
					case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
						val, _ := strconv.ParseUint(cell, 10, 64)
						newStruct.Field(k).SetUint(val)
					case reflect.Bool:
						if strings.ToUpper(cell) == "TRUE" {
							newStruct.Field(k).SetBool(true)
						} else {
							newStruct.Field(k).SetBool(false)
						}
					default:
						newStruct.Field(k).SetString(cell)
					}
					break
				}
			}
		}

		slice = reflect.Append(slice, newStruct)
	}

	values.Elem().Set(slice)
	return nil
}
