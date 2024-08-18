package presetfunc

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type DataType string

var (
	DataTypeInterface DataType = "Interface"
)

type DataTypeDetail struct {
	DataType       string
	IsSelector     bool
	IsArray        bool
	IsPointer      bool
	IsMap          bool
	HasOmitempty   bool
	MapKeyDataType string
}

type JsFieldInfo struct {
	Name string
	DataTypeDetail
}

type JsStructInfo struct {
	Name   string
	Field  []JsFieldInfo
	Extend []string
}

func GoStructToJSConverter(path string, n ast.Node) (resp *JsStructInfo, next bool) {
	next = true

	// 仅处理类型声明
	ts, ok := n.(*ast.TypeSpec)
	if !ok {
		return
	}

	// 仅处理结构体类型
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return
	}

	resp = &JsStructInfo{}
	resp.Name = ts.Name.Name

	for _, field := range st.Fields.List {
		jsFieldInfo := JsFieldInfo{}

		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			// 处理嵌套结构体或匿名字段
			fieldName = field.Type.(*ast.Ident).Name
		}

		jsonTag := ""
		if field.Tag != nil {
			tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1]) // 去除反引号
			jsonTag = tag.Get("json")
		}

		// 处理 JSON 标签为 "-"
		if jsonTag == "-" {
			continue
		}

		// 如果 JSON 标签为空或未指定字段名，使用字段名作为 JSON 标签
		if jsonTag == "" || strings.Split(jsonTag, ",")[0] == "" {
			jsonTag = fieldName
		}
		jsFieldInfo.Name = jsonTag

		parts := strings.SplitN(jsonTag, ",", 2)
		if len(parts) > 1 {
			// 检查是否有 omitempty
			options := strings.Split(parts[1], ",")
			for _, option := range options {
				if option == "omitempty" {
					jsFieldInfo.HasOmitempty = true
				}
			}
		}

		jsFieldInfo.DataTypeDetail = getFieldTypeName(field.Type, jsFieldInfo.DataTypeDetail)

		resp.Field = append(resp.Field, jsFieldInfo)
	}

	return
}

// 获取字段类型的名称
func getFieldTypeName(expr ast.Expr, dataTypeDetail DataTypeDetail) DataTypeDetail {
	switch t := expr.(type) {
	case *ast.Ident:
		dataTypeDetail.DataType = t.Name
	case *ast.StarExpr:
		dataTypeDetail.IsPointer = true
		getFieldTypeName(t.X, dataTypeDetail)
	case *ast.SelectorExpr:
		dataTypeDetail.IsSelector = true
		dataTypeDetail.DataType = t.Sel.Name
	case *ast.ArrayType:
		dataTypeDetail.IsArray = true
		getFieldTypeName(t.Elt, dataTypeDetail)
	case *ast.MapType:
		dataTypeDetail.IsMap = true
		dataTypeDetail.MapKeyDataType = getFieldTypeName(t.Key, DataTypeDetail{}).DataType
		dataTypeDetail.DataType = getFieldTypeName(t.Value, DataTypeDetail{}).DataType
	case *ast.InterfaceType:
		dataTypeDetail.DataType = string(DataTypeInterface)
	default:
		dataTypeDetail.DataType = fmt.Sprintf("%T", expr) // Fallback to type name
	}
	return dataTypeDetail
}
