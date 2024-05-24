package csvhelper

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

type lCsvColumnInfo struct {
	Index          int64
	FieldName      string
	Sort           int64
	Header         string
	Dictionary     string
	HideFromExport bool
	Converter      []string // this will run first before FieldNameConverter
}

type CsvConverter struct {
	ConvertFunc func(structVal reflect.Value, fieldName string, fieldVal reflect.Value, converterParams interface{}) reflect.Value
	Value       interface{}
}

type CsvSetting struct {
	RequireHeader bool
	IgnoreFields  []string
	Converters    map[string]*CsvConverter
}

func tagToColumnInfo(field reflect.StructField, ignoreFields []string) *lCsvColumnInfo {
	csvColumnInfo := lCsvColumnInfo{
		FieldName: field.Name,
		Header:    field.Name,
	}

	if arrayhelper.Includes(ignoreFields, csvColumnInfo.FieldName) {
		return nil
	}

	tagInfos := strings.Split(field.Tag.Get(CsvTag), ";")
	for _, tagInfo := range tagInfos {
		if strings.HasPrefix(tagInfo, CsvHide) {
			if strings.TrimPrefix(tagInfo, CsvHide) == "true" {
				return nil
			}
		}
		if strings.HasPrefix(tagInfo, CsvHeader) {
			csvColumnInfo.Header = strings.TrimPrefix(tagInfo, CsvHeader)
		}
		if strings.HasPrefix(tagInfo, CsvSort) {
			csvColumnInfo.Sort, _ = strconv.ParseInt(strings.TrimPrefix(tagInfo, CsvSort), 10, 64)
		}
		if strings.HasPrefix(tagInfo, CsvDefaultConverter) {
			csvColumnInfo.Converter = strings.Split(strings.TrimPrefix(tagInfo, CsvDefaultConverter), ",")
		}
	}

	return &csvColumnInfo
}

func StructsToCsv(
	writer *csv.Writer,
	values reflect.Value,
	csvSetting CsvSetting,
) (err error) {
	if values.Kind() != reflect.Ptr || values.Elem().Kind() != reflect.Slice {
		err = ErrNotArray
		return
	}
	if values.Elem().Len() == 0 {
		err = ErrEmptyArray
		return
	}

	elem := values.Elem().Index(0)

	if elem.Kind() == reflect.Pointer {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		err = ErrValNotStruct
		return
	}

	csvColumnInfos := []lCsvColumnInfo{}
	// 遍历结构体的每个字段
	for j := 0; j < elem.NumField(); j++ {
		field := elem.Type().Field(j) // 获取字段信息

		csvColumnInfo := tagToColumnInfo(field, csvSetting.IgnoreFields)
		if csvColumnInfo != nil {
			csvColumnInfos = append(csvColumnInfos, *csvColumnInfo)
		}
	}

	// sort csvColumnInfos, 0 to last
	sort.Slice(csvColumnInfos, func(i, j int) bool {
		if csvColumnInfos[i].Sort == 0 {
			return false
		}
		if csvColumnInfos[j].Sort == 0 {
			return true
		}
		return csvColumnInfos[i].Sort < csvColumnInfos[j].Sort
	})

	// input header
	if csvSetting.RequireHeader {
		var row []string
		for i := range csvColumnInfos {
			row = append(row, csvColumnInfos[i].Header)
		}
		writer.Write(row)
	}

	// input content
	for i := 0; i < values.Elem().Len(); i++ {
		// 获取切片中每个元素的反射值对象
		elem := values.Elem().Index(i)
		if elem.Kind() == reflect.Pointer {
			elem = elem.Elem()
		}
		var row []string
		// 确保元素是一个结构体
		// 遍历结构体的每个字段
		for j := 0; j < len(csvColumnInfos); j++ {
			fieldValue := elem.FieldByName(csvColumnInfos[j].FieldName) // 获取字段值

			for _, converter := range csvColumnInfos[j].Converter {
				fieldValue = csvSetting.Converters[converter].
					ConvertFunc(
						elem,
						csvColumnInfos[j].FieldName,
						fieldValue,
						csvSetting.Converters[converter].Value,
					)
			}
			if csvSetting.Converters[csvColumnInfos[j].FieldName] != nil {
				fieldValue = csvSetting.Converters[csvColumnInfos[j].FieldName].
					ConvertFunc(
						elem,
						csvColumnInfos[j].FieldName,
						fieldValue,
						csvSetting.Converters[csvColumnInfos[j].FieldName].Value,
					)
			}
			row = append(row, fmt.Sprintf("%v", fieldValue.Interface()))
		}
		writer.Write(row)
	}
	writer.Flush()
	if err = writer.Error(); err != nil {
		return
	}
	return
}
