package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func testUniqueSlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestUnique(t *testing.T) {
	tests := []struct {
		arr      interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 2, 3, 4, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1}},
		{[]int{}, []int{}},

		// 测试用例 string 类型
		{[]string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		{[]string{"x", "x", "x"}, []string{"x"}},
		{[]string{}, []string{}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 1.1, 3.3}, []float64{1.1, 2.2, 3.3}},
		{[]float64{4.4, 4.4, 4.4}, []float64{4.4}},
		{[]float64{}, []float64{}},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result := arrayhelper.Unique(arr)
			expected := tt.expected.([]int)
			if !testUniqueSlicesEqual(result, expected) {
				t.Errorf("Unique(%v) = %v; want %v", arr, result, expected)
			}
		case []string:
			result := arrayhelper.Unique(arr)
			expected := tt.expected.([]string)
			if !testUniqueSlicesEqual(result, expected) {
				t.Errorf("Unique(%v) = %v; want %v", arr, result, expected)
			}
		case []float64:
			result := arrayhelper.Unique(arr)
			expected := tt.expected.([]float64)
			if !testUniqueSlicesEqual(result, expected) {
				t.Errorf("Unique(%v) = %v; want %v", arr, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
