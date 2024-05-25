package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestFill(t *testing.T) {
	tests := []struct {
		arr      interface{}
		value    interface{}
		startEnd []int
		expected interface{}
	}{
		{[]int{0, 0, 0}, 0, []int{0, 5}, []int{0, 0, 0}},
		{[]int{0, 0, 0}, 1, []int{0, 5}, []int{1, 1, 1}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{}, []int{0, 0, 0, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{2}, []int{1, 2, 0, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{2, 4}, []int{1, 2, 0, 0, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{-2}, []int{1, 2, 3, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, -1}, []int{1, 0, 0, 0, 5}},

		{[]string{"a", "b", "c", "d"}, "z", []int{}, []string{"z", "z", "z", "z"}},
		{[]string{"a", "b", "c", "d"}, "z", []int{1}, []string{"a", "z", "z", "z"}},
		{[]string{"a", "b", "c", "d"}, "z", []int{1, 3}, []string{"a", "z", "z", "d"}},
		{[]string{"a", "b", "c", "d"}, "z", []int{-3}, []string{"a", "z", "z", "z"}},
		{[]string{"a", "b", "c", "d"}, "z", []int{1, -1}, []string{"a", "z", "z", "d"}},

		{[]float64{1.1, 2.2, 3.3, 4.4}, 0.0, []int{}, []float64{0.0, 0.0, 0.0, 0.0}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 0.0, []int{2}, []float64{1.1, 2.2, 0.0, 0.0}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 0.0, []int{2, 3}, []float64{1.1, 2.2, 0.0, 4.4}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 0.0, []int{-2}, []float64{1.1, 2.2, 0.0, 0.0}},
		{[]float64{1.1, 2.2, 3.3, 4.4}, 0.0, []int{1, -1}, []float64{1.1, 0.0, 0.0, 4.4}},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			value := tt.value.(int)
			result := arrayhelper.Fill(arr, value, tt.startEnd...)
			expected := tt.expected.([]int)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Fill(%v, %v, %v) = %v; want %v", arr, value, tt.startEnd, result, expected)
			}
		case []string:
			value := tt.value.(string)
			result := arrayhelper.Fill(arr, value, tt.startEnd...)
			expected := tt.expected.([]string)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Fill(%v, %v, %v) = %v; want %v", arr, value, tt.startEnd, result, expected)
			}
		case []float64:
			value := tt.value.(float64)
			result := arrayhelper.Fill(arr, value, tt.startEnd...)
			expected := tt.expected.([]float64)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Fill(%v, %v, %v) = %v; want %v", arr, value, tt.startEnd, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestFillWithExtensible(t *testing.T) {
	tests := []struct {
		arr      []int
		value    int
		startEnd []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 9, nil, []int{9, 9, 9, 9}},
		{[]int{1, 2, 3, 4}, 9, []int{1}, []int{1, 9, 9, 9}},
		{[]int{1, 2, 3, 4}, 9, []int{1, 3}, []int{1, 9, 9, 4}},
		{[]int{1, 2, 3, 4}, 9, []int{-2, 3}, []int{1, 2, 9, 4}},
		{[]int{1, 2, 3, 4}, 9, []int{2, 6}, []int{1, 2, 9, 9, 9, 9}},
		{[]int{1, 2, 3, 4}, 9, []int{6}, []int{1, 2, 3, 4, 0, 0, 9}},
		{[]int{1, 2, 3, 4}, 0, []int{1, -1}, []int{1, 0, 0, 4}},
	}

	for _, tt := range tests {
		result := arrayhelper.FillWithExtensible(tt.arr, tt.value, tt.startEnd...)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("FillWithExtensible(%v, %v, %v) = %v; want %v", tt.arr, tt.value, tt.startEnd, result, tt.expected)
		}
	}
}
