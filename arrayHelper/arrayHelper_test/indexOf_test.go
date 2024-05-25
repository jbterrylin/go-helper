package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestIndexOf(t *testing.T) {
	tests := []struct {
		slice         interface{}
		searchElement interface{}
		fromIndex     []int
		expected      int
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, 3, []int{}, 2},
		{[]int{1, 2, 3, 4, 5}, 6, []int{}, -1},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2}, 2},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3}, -1},
		{[]int{1, 2, 3, 4, 5}, 3, []int{-3}, 2},
		{[]int{1, 2, 3, 4, 5}, 3, []int{-1}, -1},

		// 测试用例 string 类型
		{[]string{"a", "b", "c", "d"}, "c", []int{}, 2},
		{[]string{"a", "b", "c", "d"}, "e", []int{}, -1},
		{[]string{"a", "b", "c", "d"}, "c", []int{2}, 2},
		{[]string{"a", "b", "c", "d"}, "c", []int{3}, -1},
		{[]string{"a", "b", "c", "d"}, "c", []int{-2}, 2},
		{[]string{"a", "b", "c", "d"}, "c", []int{-1}, -1},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{}, 2},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 5.5, []int{}, -1},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{2}, 2},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{3}, -1},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{-2}, 2},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3.3, []int{-1}, -1},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case []int:
			result := arrayhelper.IndexOf(slice, tt.searchElement.(int), tt.fromIndex...)
			if result != tt.expected {
				t.Errorf("IndexOf(%v, %v, %v) = %v; want %v", slice, tt.searchElement, tt.fromIndex, result, tt.expected)
			}
		case []string:
			result := arrayhelper.IndexOf(slice, tt.searchElement.(string), tt.fromIndex...)
			if result != tt.expected {
				t.Errorf("IndexOf(%v, %v, %v) = %v; want %v", slice, tt.searchElement, tt.fromIndex, result, tt.expected)
			}
		case []float64:
			result := arrayhelper.IndexOf(slice, tt.searchElement.(float64), tt.fromIndex...)
			if result != tt.expected {
				t.Errorf("IndexOf(%v, %v, %v) = %v; want %v", slice, tt.searchElement, tt.fromIndex, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
