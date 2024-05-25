package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestKeys(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, []int{0, 1, 2}},
		{[]int{5, 6, 7, 8}, []int{0, 1, 2, 3}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, []int{0, 1, 2}},
		{[]string{"x", "y"}, []int{0, 1}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3}, []int{0, 1, 2}},
		{[]float64{4.4, 5.5}, []int{0, 1}},
	}

	for _, tt := range tests {
		switch input := tt.input.(type) {
		case []int:
			result := arrayhelper.Keys(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Keys(%v) = %v; want %v", input, result, tt.expected)
			}
		case []string:
			result := arrayhelper.Keys(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Keys(%v) = %v; want %v", input, result, tt.expected)
			}
		case []float64:
			result := arrayhelper.Keys(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Keys(%v) = %v; want %v", input, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.input)
		}
	}
}
