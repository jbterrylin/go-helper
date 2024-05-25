package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestIterator(t *testing.T) {
	tests := []struct {
		slice    interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, []arrayhelper.Entry[int]{{Index: 0, Value: 1}, {Index: 1, Value: 2}, {Index: 2, Value: 3}}},
		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, []arrayhelper.Entry[string]{{Index: 0, Value: "a"}, {Index: 1, Value: "b"}, {Index: 2, Value: "c"}}},
		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3}, []arrayhelper.Entry[float64]{{Index: 0, Value: 1.1}, {Index: 1, Value: 2.2}, {Index: 2, Value: 3.3}}},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case []int:
			it := arrayhelper.NewIterator(slice)
			var result []arrayhelper.Entry[int]
			for {
				entry, ok := it.Next()
				if !ok {
					break
				}
				result = append(result, entry)
			}
			expected := tt.expected.([]arrayhelper.Entry[int])
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Iterator(%v) = %v; want %v", slice, result, expected)
			}
		case []string:
			it := arrayhelper.NewIterator(slice)
			var result []arrayhelper.Entry[string]
			for {
				entry, ok := it.Next()
				if !ok {
					break
				}
				result = append(result, entry)
			}
			expected := tt.expected.([]arrayhelper.Entry[string])
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Iterator(%v) = %v; want %v", slice, result, expected)
			}
		case []float64:
			it := arrayhelper.NewIterator(slice)
			var result []arrayhelper.Entry[float64]
			for {
				entry, ok := it.Next()
				if !ok {
					break
				}
				result = append(result, entry)
			}
			expected := tt.expected.([]arrayhelper.Entry[float64])
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Iterator(%v) = %v; want %v", slice, result, expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
