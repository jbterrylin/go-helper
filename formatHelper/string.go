package formathelper

import (
	"strings"
	"unicode"
)

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// ToCamelCase 将字符串转换为 CamelCase
func ToCamelCase(s string) string {
	if s == "" {
		return s
	}
	parts := splitOnNonAlnum(s)
	for i := range parts {
		parts[i] = strings.Title(strings.ToLower(parts[i]))
	}
	return strings.Join(parts, "")
}

// ToKebabCase 将字符串转换为 kebab-case
func ToKebabCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(strings.Join(splitOnNonAlnum(s), "-"))
}

// ToSnakeCase 将字符串转换为 snake_case
func ToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(strings.Join(splitOnNonAlnum(s), "_"))
}

// ToLowerCamelCase 将字符串转换为 lowerCamelCase
func ToLowerCamelCase(s string) string {
	if s == "" {
		return s
	}
	camel := ToCamelCase(s)
	return strings.ToLower(string(camel[0])) + camel[1:]
}

// ToUpperCamelCase 将字符串转换为 UpperCamelCase
func ToUpperCamelCase(s string) string {
	if s == "" {
		return s
	}
	return ToCamelCase(s)
}

// ToPascalCase 将字符串转换为 PascalCase（与 UpperCamelCase 相同）
func ToPascalCase(s string) string {
	if s == "" {
		return s
	}
	return ToCamelCase(s)
}

// ToTitleCase 将字符串转换为 Title Case
func ToTitleCase(s string) string {
	if s == "" {
		return s
	}
	return strings.Title(strings.ToLower(s))
}

// ToSpaceCase 将字符串转换为空格分隔
func ToSpaceCase(s string) string {
	if s == "" {
		return s
	}
	return strings.Join(splitOnNonAlnum(s), " ")
}

// ToTabCase 将字符串转换为制表符分隔
func ToTabCase(s string) string {
	if s == "" {
		return s
	}
	return strings.Join(splitOnNonAlnum(s), "\t")
}

// splitOnNonAlnum 将字符串按非字母数字字符分割
func splitOnNonAlnum(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
