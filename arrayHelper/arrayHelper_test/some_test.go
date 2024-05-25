package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestSome(t *testing.T) {
	tests := []struct {
		slice    interface{}
		callback interface{}
		expected bool
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v > 3 }, true},
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v > 5 }, false},
		{[]int{}, func(v int) bool { return v > 0 }, false},

		// 测试用例 string 类型
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "banana" }, true},
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "date" }, false},
		{[]string{}, func(v string) bool { return v == "anything" }, false},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v > 2.2 }, true},
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v > 5.5 }, false},
		{[]float64{}, func(v float64) bool { return v > 0 }, false},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case []int:
			result := arrayhelper.Some(slice, tt.callback.(func(int) bool))
			if result != tt.expected {
				t.Errorf("Some(%v, callback) = %v; want %v", slice, result, tt.expected)
			}
		case []string:
			result := arrayhelper.Some(slice, tt.callback.(func(string) bool))
			if result != tt.expected {
				t.Errorf("Some(%v, callback) = %v; want %v", slice, result, tt.expected)
			}
		case []float64:
			result := arrayhelper.Some(slice, tt.callback.(func(float64) bool))
			if result != tt.expected {
				t.Errorf("Some(%v, callback) = %v; want %v", slice, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
