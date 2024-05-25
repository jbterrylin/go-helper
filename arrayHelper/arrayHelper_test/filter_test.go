package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		arr      interface{}
		fn       interface{}
		expected interface{}
	}{
		// Test cases for int
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v > 3 }, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, func(v int) bool { return v < 0 }, []int{}},

		// Test cases for string
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return len(v) > 5 }, []string{"banana", "cherry"}},
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v[0] == 'a' }, []string{"apple"}},
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return v == "grape" }, []string{}},

		// Test cases for float64
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v > 2.0 }, []float64{2.2, 3.3, 4.4}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v < 1.5 }, []float64{1.1}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, func(v float64) bool { return v == 5.5 }, []float64{}},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			fn := tt.fn.(func(int) bool)
			result := arrayhelper.Filter(arr, fn)
			expected := tt.expected.([]int)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Filter(%v, fn) = %v; want %v", arr, result, expected)
			}
		case []string:
			fn := tt.fn.(func(string) bool)
			result := arrayhelper.Filter(arr, fn)
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Filter(%v, fn) = %v; want %v", arr, result, expected)
			}
		case []float64:
			fn := tt.fn.(func(float64) bool)
			result := arrayhelper.Filter(arr, fn)
			expected := tt.expected.([]float64)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Filter(%v, fn) = %v; want %v", arr, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
