package arrayhelper_test

import (
	"errors"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestFind(t *testing.T) {
	tests := []struct {
		arr      interface{}
		test     interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v == 3 }, 3},
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v == 6 }, 0},

		// 测试用例 string 类型
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "banana" }, "banana"},
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "grape" }, ""},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 2.2 }, 2.2},
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 5.5 }, 0.0},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			fn := tt.test.(func(int) bool)
			result := arrayhelper.Find(arr, fn)
			expected := tt.expected.(int)
			if result != expected {
				t.Errorf("Find(%v, fn) = %v; want %v", arr, result, expected)
			}
		case []string:
			fn := tt.test.(func(string) bool)
			result := arrayhelper.Find(arr, fn)
			expected := tt.expected.(string)
			if result != expected {
				t.Errorf("Find(%v, fn) = %v; want %v", arr, result, expected)
			}
		case []float64:
			fn := tt.test.(func(float64) bool)
			result := arrayhelper.Find(arr, fn)
			expected := tt.expected.(float64)
			if result != expected {
				t.Errorf("Find(%v, fn) = %v; want %v", arr, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestFindNth(t *testing.T) {
	tests := []struct {
		arr      []int
		test     func(int) bool
		findNth  int
		expected int
	}{
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 3 }, 2, 3},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 3 }, 3, 0},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 2 }, 1, 2},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 2 }, 2, 0},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 5 }, 1, 5},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 6 }, 1, 0},
	}

	for _, tt := range tests {
		result := arrayhelper.FindNth(tt.arr, tt.test, tt.findNth)
		if result != tt.expected {
			t.Errorf("FindNth(%v, test, %d) = %v; want %v", tt.arr, tt.findNth, result, tt.expected)
		}
	}
}

func TestFindWithErr(t *testing.T) {
	tests := []struct {
		arr      interface{}
		test     interface{}
		expected interface{}
		err      error
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v == 3 }, 3, nil},
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v == 6 }, 0, arrayhelper.ErrNotFound},

		// 测试用例 string 类型
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "banana" }, "banana", nil},
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "grape" }, "", arrayhelper.ErrNotFound},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 2.2 }, 2.2, nil},
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 5.5 }, 0.0, arrayhelper.ErrNotFound},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			fn := tt.test.(func(int) bool)
			result, err := arrayhelper.FindWithErr(arr, fn)
			expected := tt.expected.(int)
			if result != expected || !errors.Is(err, tt.err) {
				t.Errorf("FindWithErr(%v, fn) = %v, %v; want %v, %v", arr, result, err, expected, tt.err)
			}
		case []string:
			fn := tt.test.(func(string) bool)
			result, err := arrayhelper.FindWithErr(arr, fn)
			expected := tt.expected.(string)
			if result != expected || !errors.Is(err, tt.err) {
				t.Errorf("FindWithErr(%v, fn) = %v, %v; want %v, %v", arr, result, err, expected, tt.err)
			}
		case []float64:
			fn := tt.test.(func(float64) bool)
			result, err := arrayhelper.FindWithErr(arr, fn)
			expected := tt.expected.(float64)
			if result != expected || !errors.Is(err, tt.err) {
				t.Errorf("FindWithErr(%v, fn) = %v, %v; want %v, %v", arr, result, err, expected, tt.err)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestFindNthWithErr(t *testing.T) {
	tests := []struct {
		arr      []int
		test     func(int) bool
		findNth  int
		expected int
		err      error
	}{
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 3 }, 2, 3, nil},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 3 }, 3, 0, arrayhelper.ErrNotFound},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 2 }, 1, 2, nil},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 2 }, 2, 0, arrayhelper.ErrNotFound},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 5 }, 1, 5, nil},
		{[]int{1, 2, 3, 3, 4, 5}, func(v int) bool { return v == 6 }, 1, 0, arrayhelper.ErrNotFound},
	}

	for _, tt := range tests {
		result, err := arrayhelper.FindNthWithErr(tt.arr, tt.test, tt.findNth)
		if !errors.Is(err, tt.err) || result != tt.expected {
			t.Errorf("FindNthWithErr(%v, test, %d) = %v, %v; want %v, %v", tt.arr, tt.findNth, result, err, tt.expected, tt.err)
		}
	}
}
