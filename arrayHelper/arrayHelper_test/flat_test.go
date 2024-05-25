package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][][]int{{{1, 2}, {3, 4}}, {{5, 6}}}, []int{1, 2, 3, 4, 5, 6}},
		{[]interface{}{[]int{1, 2}, []interface{}{3, []int{4, 5}}}, []int{1, 2, 3, 4, 5}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[][]string{{"a", "b"}, {"c", "d"}}, []string{"a", "b", "c", "d"}},
		{[][][]string{{{"a", "b"}, {"c", "d"}}, {{"e", "f"}}}, []string{"a", "b", "c", "d", "e", "f"}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}},
		{[][]float64{{1.1, 2.2}, {3.3, 4.4}}, []float64{1.1, 2.2, 3.3, 4.4}},
		{[][][]float64{{{1.1, 2.2}, {3.3, 4.4}}, {{5.5, 6.6}}}, []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}},
	}

	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case []int:
			result := arrayhelper.Flatten[int](tt.input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Flatten(%v) = %v; want %v", tt.input, result, expected)
			}
		case []string:
			result := arrayhelper.Flatten[string](tt.input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Flatten(%v) = %v; want %v", tt.input, result, expected)
			}
		case []float64:
			result := arrayhelper.Flatten[float64](tt.input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Flatten(%v) = %v; want %v", tt.input, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.expected)
		}
	}
}

func TestFlattenAll(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][][]int{{{1, 2}, {3, 4}}, {{5, 6}}}, []int{1, 2, 3, 4, 5, 6}},
		{[]interface{}{[]int{1, 2}, []interface{}{3, []int{4, 5}}}, []int{1, 2, 3, 4, 5}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[][]string{{"a", "b"}, {"c", "d"}}, []string{"a", "b", "c", "d"}},
		{[][][]string{{{"a", "b"}, {"c", "d"}}, {{"e", "f"}}}, []string{"a", "b", "c", "d", "e", "f"}},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}},
		{[][]float64{{1.1, 2.2}, {3.3, 4.4}}, []float64{1.1, 2.2, 3.3, 4.4}},
		{[][][]float64{{{1.1, 2.2}, {3.3, 4.4}}, {{5.5, 6.6}}}, []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}},
	}

	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case []int:
			result := arrayhelper.FlattenAll[int](tt.input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("FlattenAll(%v) = %v; want %v", tt.input, result, expected)
			}
		case []string:
			result := arrayhelper.FlattenAll[string](tt.input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("FlattenAll(%v) = %v; want %v", tt.input, result, expected)
			}
		case []float64:
			result := arrayhelper.FlattenAll[float64](tt.input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("FlattenAll(%v) = %v; want %v", tt.input, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.expected)
		}
	}
}
