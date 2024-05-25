package arrayhelper_test

import (
	"fmt"
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestInit(t *testing.T) {
	tests := []struct {
		length   int
		setValue func(int) interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{5, func(i int) interface{} { return i * 2 }, []int{0, 2, 4, 6, 8}},
		{3, func(i int) interface{} { return i + 1 }, []int{1, 2, 3}},

		// 测试用例 string 类型
		{4, func(i int) interface{} { return fmt.Sprintf("%c", 'a'+i) }, []string{"a", "b", "c", "d"}},

		// 测试用例 float64 类型
		{3, func(i int) interface{} { return float64(i) * 1.1 }, []float64{0, 1.1, 2.2}},
	}

	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case []int:
			result := arrayhelper.Init(tt.length, func(i int) int { return tt.setValue(i).(int) })
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Init(%v, setValue) = %v; want %v", tt.length, result, expected)
			}
		case []string:
			result := arrayhelper.Init(tt.length, func(i int) string { return tt.setValue(i).(string) })
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Init(%v, setValue) = %v; want %v", tt.length, result, expected)
			}
		case []float64:
			result := arrayhelper.Init(tt.length, func(i int) float64 { return tt.setValue(i).(float64) })
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Init(%v, setValue) = %v; want %v", tt.length, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.expected)
		}
	}
}

func TestInitConst(t *testing.T) {
	tests := []struct {
		length   int
		value    interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{5, 42, []int{42, 42, 42, 42, 42}},
		{3, 7, []int{7, 7, 7}},

		// 测试用例 string 类型
		{4, "hello", []string{"hello", "hello", "hello", "hello"}},

		// 测试用例 float64 类型
		{3, 3.14, []float64{3.14, 3.14, 3.14}},
	}

	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case []int:
			result := arrayhelper.InitConst(tt.length, tt.value.(int))
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("InitConst(%v, %v) = %v; want %v", tt.length, tt.value, result, expected)
			}
		case []string:
			result := arrayhelper.InitConst(tt.length, tt.value.(string))
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("InitConst(%v, %v) = %v; want %v", tt.length, tt.value, result, expected)
			}
		case []float64:
			result := arrayhelper.InitConst(tt.length, tt.value.(float64))
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("InitConst(%v, %v) = %v; want %v", tt.length, tt.value, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.expected)
		}
	}
}
