package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3}, []float64{3.3, 2.2, 1.1}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, []float64{4.4, 3.3, 2.2, 1.1}},
	}

	for _, tt := range tests {
		switch input := tt.input.(type) {
		case []int:
			arrayhelper.Reverse(input)
			expected := tt.expected.([]int)
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Reverse(%v) = %v; want %v", tt.input, input, expected)
			}
		case []string:
			arrayhelper.Reverse(input)
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Reverse(%v) = %v; want %v", tt.input, input, expected)
			}
		case []float64:
			arrayhelper.Reverse(input)
			expected := tt.expected.([]float64)
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Reverse(%v) = %v; want %v", tt.input, input, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.input)
		}
	}
}
