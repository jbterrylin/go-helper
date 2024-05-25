package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestIncludes(t *testing.T) {
	tests := []struct {
		slice         interface{}
		searchElement interface{}
		fromIndex     []int
		expected      bool
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, 3, []int{}, true},
		{[]int{1, 2, 3, 4, 5}, 6, []int{}, false},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2}, true},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3}, false},
		{[]int{1, 2, 3, 4, 5}, 3, []int{-3}, true},
		{[]int{1, 2, 3, 4, 5}, 3, []int{-1}, false},

		// 测试用例 string 类型
		{[]string{"a", "b", "c", "d"}, "c", []int{}, true},
		{[]string{"a", "b", "c", "d"}, "e", []int{}, false},
		{[]string{"a", "b", "c", "d"}, "c", []int{2}, true},
		{[]string{"a", "b", "c", "d"}, "c", []int{3}, false},
		{[]string{"a", "b", "c", "d"}, "c", []int{-2}, true},
		{[]string{"a", "b", "c", "d"}, "c", []int{-1}, false},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{}, true},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 5.5, []int{}, false},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{2}, true},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{3}, false},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{-2}, true},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{-1}, false},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case []int:
			result := arrayhelper.Includes(slice, tt.searchElement.(int), tt.fromIndex...)
			if result != tt.expected {
				t.Errorf("Includes(%v, %v, %v) = %v; want %v", slice, tt.searchElement, tt.fromIndex, result, tt.expected)
			}
		case []string:
			result := arrayhelper.Includes(slice, tt.searchElement.(string), tt.fromIndex...)
			if result != tt.expected {
				t.Errorf("Includes(%v, %v, %v) = %v; want %v", slice, tt.searchElement, tt.fromIndex, result, tt.expected)
			}
		case []float64:
			result := arrayhelper.Includes(slice, tt.searchElement.(float64), tt.fromIndex...)
			if result != tt.expected {
				t.Errorf("Includes(%v, %v, %v) = %v; want %v", slice, tt.searchElement, tt.fromIndex, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
