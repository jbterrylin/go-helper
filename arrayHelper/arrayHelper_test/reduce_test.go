package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		slice    interface{}
		reducer  interface{}
		initial  interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3, 4}, func(acc int, v int) int { return acc + v }, 0, 10},
		{[]int{1, 2, 3, 4}, func(acc int, v int) int { return acc * v }, 1, 24},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, func(acc string, v string) string { return acc + v }, "", "abc"},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case []int:
			switch reducer := tt.reducer.(type) {
			case func(int, int) int:
				result := arrayhelper.Reduce(slice, reducer, tt.initial.(int))
				if result != tt.expected {
					t.Errorf("Reduce(%v, reducer, %v) = %v; want %v", slice, tt.initial, result, tt.expected)
				}
			}
		case []string:
			switch reducer := tt.reducer.(type) {
			case func(string, string) string:
				result := arrayhelper.Reduce(slice, reducer, tt.initial.(string))
				if result != tt.expected {
					t.Errorf("Reduce(%v, reducer, %v) = %v; want %v", slice, tt.initial, result, tt.expected)
				}
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
