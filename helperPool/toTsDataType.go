package helperpool

import "strings"

func GoTypeToTsType(goType string) string {
	goType = strings.TrimSpace(goType)

	switch goType {
	case "string":
		return "string"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "[]byte":
		return "Uint8Array"
	case "map[string]interface{}", "map[string]any":
		return "{ [key: string]: any }"
	case "interface{}", "any":
		return "any"
	default:
		// 默认情况下返回原始类型
		return goType
	}
}
