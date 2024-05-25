package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		arr      interface{}
		times    int
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, -2, []int{3, 4, 5, 1, 2}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c", "d"}, 1, []string{"d", "a", "b", "c"}},
		{[]string{"a", "b", "c", "d"}, -1, []string{"b", "c", "d", "a"}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, 2, []float64{3.3, 4.4, 1.1, 2.2}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, -2, []float64{3.3, 4.4, 1.1, 2.2}},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result := arrayhelper.Rotate(arr, tt.times)
			expected := tt.expected.([]int)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Rotate(%v, %d) = %v; want %v", arr, tt.times, result, expected)
			}
		case []string:
			result := arrayhelper.Rotate(arr, tt.times)
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Rotate(%v, %d) = %v; want %v", arr, tt.times, result, expected)
			}
		case []float64:
			result := arrayhelper.Rotate(arr, tt.times)
			expected := tt.expected.([]float64)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Rotate(%v, %d) = %v; want %v", arr, tt.times, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestRotateByIndex(t *testing.T) {
	tests := []struct {
		arr      interface{}
		index    int
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, -2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},  // invalid index
		{[]int{1, 2, 3, 4, 5}, -6, []int{1, 2, 3, 4, 5}}, // invalid index

		// 测试用例 string 类型
		{[]string{"a", "b", "c", "d"}, 1, []string{"b", "c", "d", "a"}},
		{[]string{"a", "b", "c", "d"}, -1, []string{"d", "a", "b", "c"}},
		{[]string{"a", "b", "c", "d"}, 4, []string{"a", "b", "c", "d"}},  // invalid index
		{[]string{"a", "b", "c", "d"}, -5, []string{"a", "b", "c", "d"}}, // invalid index

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, 2, []float64{3.3, 4.4, 1.1, 2.2}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, -2, []float64{3.3, 4.4, 1.1, 2.2}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 4, []float64{1.1, 2.2, 3.3, 4.4}},  // invalid index
		{[]float64{1.1, 2.2, 3.3, 4.4}, -5, []float64{1.1, 2.2, 3.3, 4.4}}, // invalid index
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result := arrayhelper.RotateByIndex(arr, tt.index)
			expected := tt.expected.([]int)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("RotateByIndex(%v, %d) = %v; want %v", arr, tt.index, result, expected)
			}
		case []string:
			result := arrayhelper.RotateByIndex(arr, tt.index)
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("RotateByIndex(%v, %d) = %v; want %v", arr, tt.index, result, expected)
			}
		case []float64:
			result := arrayhelper.RotateByIndex(arr, tt.index)
			expected := tt.expected.([]float64)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("RotateByIndex(%v, %d) = %v; want %v", arr, tt.index, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestRotateByIndexSafe(t *testing.T) {
	tests := []struct {
		arr      []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, -2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{5, 1, 2, 3, 4}},  // invalid index, should use closest valid
		{[]int{1, 2, 3, 4, 5}, -6, []int{1, 2, 3, 4, 5}}, // invalid index, should use closest valid
	}

	for _, tt := range tests {
		result := arrayhelper.RotateByIndexSafe(tt.arr, tt.index)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("RotateByIndexSafe(%v, %d) = %v; want %v", tt.arr, tt.index, result, tt.expected)
		}
	}
}
