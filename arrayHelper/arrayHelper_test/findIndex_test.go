package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestFindIndex(t *testing.T) {
	tests := []struct {
		arr      interface{}
		test     interface{}
		expected int
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v == 3 }, 2},
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v == 6 }, -1},

		// 测试用例 string 类型
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "banana" }, 1},
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "grape" }, -1},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 2.2 }, 1},
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 5.5 }, -1},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			fn := tt.test.(func(int) bool)
			result := arrayhelper.FindIndex(arr, fn)
			if result != tt.expected {
				t.Errorf("FindIndex(%v, fn) = %v; want %v", arr, result, tt.expected)
			}
		case []string:
			fn := tt.test.(func(string) bool)
			result := arrayhelper.FindIndex(arr, fn)
			if result != tt.expected {
				t.Errorf("FindIndex(%v, fn) = %v; want %v", arr, result, tt.expected)
			}
		case []float64:
			fn := tt.test.(func(float64) bool)
			result := arrayhelper.FindIndex(arr, fn)
			if result != tt.expected {
				t.Errorf("FindIndex(%v, fn) = %v; want %v", arr, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestFindNthIndex(t *testing.T) {
	tests := []struct {
		arr      interface{}
		test     interface{}
		findNth  int
		force    bool
		expected int
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5, 2}, func(v int) bool { return v == 2 }, 2, false, 5},
		{[]int{1, 2, 3, 4, 5, 2}, func(v int) bool { return v == 2 }, 3, false, -1},
		{[]int{1, 2, 3, 4, 5, 2}, func(v int) bool { return v == 2 }, 3, true, -1},

		// 测试用例 string 类型
		{[]string{"apple", "banana", "cherry", "banana"}, func(v string) bool { return v == "banana" }, 2, false, 3},
		{[]string{"apple", "banana", "cherry", "banana"}, func(v string) bool { return v == "banana" }, 3, false, -1},
		{[]string{"apple", "banana", "cherry", "banana"}, func(v string) bool { return v == "banana" }, 3, true, -1},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4, 2.2}, func(v float64) bool { return v == 2.2 }, 2, false, 4},
		{[]float64{1.1, 2.2, 3.3, 4.4, 2.2}, func(v float64) bool { return v == 2.2 }, 3, false, -1},
		{[]float64{1.1, 2.2, 3.3, 4.4, 2.2}, func(v float64) bool { return v == 2.2 }, 3, true, -1},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			fn := tt.test.(func(int) bool)
			result := arrayhelper.FindNthIndex(arr, fn, tt.findNth, tt.force)
			if result != tt.expected {
				t.Errorf("FindNthIndex(%v, fn, %d, %v) = %v; want %v", arr, tt.findNth, tt.force, result, tt.expected)
			}
		case []string:
			fn := tt.test.(func(string) bool)
			result := arrayhelper.FindNthIndex(arr, fn, tt.findNth, tt.force)
			if result != tt.expected {
				t.Errorf("FindNthIndex(%v, fn, %d, %v) = %v; want %v", arr, tt.findNth, tt.force, result, tt.expected)
			}
		case []float64:
			fn := tt.test.(func(float64) bool)
			result := arrayhelper.FindNthIndex(arr, fn, tt.findNth, tt.force)
			if result != tt.expected {
				t.Errorf("FindNthIndex(%v, fn, %d, %v) = %v; want %v", arr, tt.findNth, tt.force, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
