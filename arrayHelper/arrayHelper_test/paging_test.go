package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestPaging(t *testing.T) {
	tests := []struct {
		arr      interface{}
		page     int
		pageSize int
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, 1, 2, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 2, 2, []int{3, 4}},
		{[]int{1, 2, 3, 4, 5}, 3, 2, []int{5}},
		{[]int{1, 2, 3, 4, 5}, 4, 2, []int{}},
		{[]int{1, 2, 3, 4, 5}, -1, 2, []int{1, 2}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c", "d"}, 1, 2, []string{"a", "b"}},
		{[]string{"a", "b", "c", "d"}, 2, 2, []string{"c", "d"}},
		{[]string{"a", "b", "c", "d"}, 3, 2, []string{}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, 1, 2, []float64{1.1, 2.2}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 2, 2, []float64{3.3, 4.4}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 3, 2, []float64{}},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result := arrayhelper.Paging(arr, tt.page, tt.pageSize)
			expected := tt.expected.([]int)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Paging(%v, %d, %d) = %v; want %v", arr, tt.page, tt.pageSize, result, expected)
			}
		case []string:
			result := arrayhelper.Paging(arr, tt.page, tt.pageSize)
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Paging(%v, %d, %d) = %v; want %v", arr, tt.page, tt.pageSize, result, expected)
			}
		case []float64:
			result := arrayhelper.Paging(arr, tt.page, tt.pageSize)
			expected := tt.expected.([]float64)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Paging(%v, %d, %d) = %v; want %v", arr, tt.page, tt.pageSize, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
