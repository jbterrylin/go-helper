package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestValueEqual(t *testing.T) {
	tests := []struct {
		firstArr  interface{}
		secondArr interface{}
		expected  bool
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{1, 2, 2}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{}, []int{}, true},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{[]string{}, []string{}, true},

		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3}, []float64{3.3, 2.2, 1.1}, true},
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 2.2}, false},
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3, 4.4}, false},
		{[]float64{}, []float64{}, true},
	}

	for _, tt := range tests {
		switch firstArr := tt.firstArr.(type) {
		case []int:
			secondArr := tt.secondArr.([]int)
			result := arrayhelper.ValueEqual(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("ValueEqual(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		case []string:
			secondArr := tt.secondArr.([]string)
			result := arrayhelper.ValueEqual(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("ValueEqual(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		case []float64:
			secondArr := tt.secondArr.([]float64)
			result := arrayhelper.ValueEqual(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("ValueEqual(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.firstArr)
		}
	}
}
