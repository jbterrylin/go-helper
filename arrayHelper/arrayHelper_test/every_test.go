package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestEvery(t *testing.T) {
	tests := []struct {
		arr      interface{}
		testFunc interface{}
		expected bool
	}{
		// Testing with int
		{[]int{2, 4, 6}, func(v int) bool { return v%2 == 0 }, true},
		{[]int{2, 3, 6}, func(v int) bool { return v%2 == 0 }, false},
		{[]int{}, func(v int) bool { return v%2 == 0 }, true},
		{[]int{1, 3, 5}, func(v int) bool { return v%2 != 0 }, true},

		// Testing with string
		{[]string{"apple", "banana", "cherry"}, func(v string) bool { return len(v) > 0 }, true},
		{[]string{"apple", "", "cherry"}, func(v string) bool { return len(v) > 0 }, false},

		// Testing with float64
		{[]float64{1.1, 2.2, 3.3}, func(v float64) bool { return v > 0 }, true},
		{[]float64{1.1, -2.2, 3.3}, func(v float64) bool { return v > 0 }, false},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			testFunc := tt.testFunc.(func(int) bool)
			result := arrayhelper.Every(arr, testFunc)
			if result != tt.expected {
				t.Errorf("Every(%v, testFunc) = %v; want %v", arr, result, tt.expected)
			}
		case []string:
			testFunc := tt.testFunc.(func(string) bool)
			result := arrayhelper.Every(arr, testFunc)
			if result != tt.expected {
				t.Errorf("Every(%v, testFunc) = %v; want %v", arr, result, tt.expected)
			}
		case []float64:
			testFunc := tt.testFunc.(func(float64) bool)
			result := arrayhelper.Every(arr, testFunc)
			if result != tt.expected {
				t.Errorf("Every(%v, testFunc) = %v; want %v", arr, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
