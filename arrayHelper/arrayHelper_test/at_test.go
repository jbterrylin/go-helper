package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestAt(t *testing.T) {
	tests := []struct {
		arr      interface{}
		index    int
		expected interface{}
	}{
		{[]int{}, 3, 0}, // invalid index
		{[]int{1, 2, 3}, 1, 2},
		{[]int{1, 2, 3}, -1, 0}, // invalid index
		{[]int{1, 2, 3}, 3, 0},  // invalid index
		{[]string{"a", "b", "c"}, 2, "c"},
		{[]string{"a", "b", "c"}, 3, ""},  // invalid index
		{[]string{"a", "b", "c"}, -1, ""}, // invalid index
		{[]float64{1.1, 2.2, 3.3}, 0, 1.1},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result := arrayhelper.At(arr, tt.index)
			if result != tt.expected {
				t.Errorf("At(%v, %d) = %v; want %v", arr, tt.index, result, tt.expected)
			}
		case []string:
			result := arrayhelper.At(arr, tt.index)
			if result != tt.expected {
				t.Errorf("At(%v, %d) = %v; want %v", arr, tt.index, result, tt.expected)
			}
		case []float64:
			result := arrayhelper.At(arr, tt.index)
			if result != tt.expected {
				t.Errorf("At(%v, %d) = %v; want %v", arr, tt.index, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
