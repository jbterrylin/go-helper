package arrayhelper_test

import (
	"fmt"
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestMap(t *testing.T) {
	tests := []struct {
		slice     interface{}
		transform interface{}
		expected  interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, func(v int) int { return v * 2 }, []int{2, 4, 6}},
		{[]int{1, 2, 3}, func(v int) string { return fmt.Sprint(v) }, []string{"1", "2", "3"}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, func(v string) string { return v + v }, []string{"aa", "bb", "cc"}},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case []int:
			switch transform := tt.transform.(type) {
			case func(int) int:
				result := arrayhelper.Map(slice, transform)
				expected := tt.expected.([]int)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Map(%v, transform) = %v; want %v", slice, result, expected)
				}
			case func(int) string:
				result := arrayhelper.Map(slice, transform)
				expected := tt.expected.([]string)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Map(%v, transform) = %v; want %v", slice, result, expected)
				}
			}
		case []string:
			result := arrayhelper.Map(slice, tt.transform.(func(string) string))
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Map(%v, transform) = %v; want %v", slice, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
