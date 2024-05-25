package arrayhelper_test

import (
	"errors"
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		arr      interface{}
		index    int
		value    interface{}
		expected interface{}
		err      error
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, 1, 9, []int{1, 9, 2, 3}, nil},
		{[]int{1, 2, 3}, 0, 9, []int{9, 1, 2, 3}, nil},
		{[]int{1, 2, 3}, 3, 9, []int{1, 2, 3, 9}, nil},
		{[]int{1, 2, 3}, -1, 9, nil, arrayhelper.ErrOutOfRange},
		{[]int{1, 2, 3}, 4, 9, nil, arrayhelper.ErrOutOfRange},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, 1, "x", []string{"a", "x", "b", "c"}, nil},
		{[]string{"a", "b", "c"}, 0, "x", []string{"x", "a", "b", "c"}, nil},
		{[]string{"a", "b", "c"}, 3, "x", []string{"a", "b", "c", "x"}, nil},
		{[]string{"a", "b", "c"}, -1, "x", nil, arrayhelper.ErrOutOfRange},
		{[]string{"a", "b", "c"}, 4, "x", nil, arrayhelper.ErrOutOfRange},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			if tt.err != nil {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Insert did not panic for index %d in %v", tt.index, arr)
					}
				}()
				arrayhelper.Insert(arr, tt.index, tt.value.(int))
			} else {
				result := arrayhelper.Insert(arr, tt.index, tt.value.(int))
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Insert(%v, %d, %v) = %v; want %v", arr, tt.index, tt.value, result, tt.expected)
				}
			}
		case []string:
			if tt.err != nil {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Insert did not panic for index %d in %v", tt.index, arr)
					}
				}()
				arrayhelper.Insert(arr, tt.index, tt.value.(string))
			} else {
				result := arrayhelper.Insert(arr, tt.index, tt.value.(string))
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Insert(%v, %d, %v) = %v; want %v", arr, tt.index, tt.value, result, tt.expected)
				}
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestInsertSafe(t *testing.T) {
	tests := []struct {
		arr      interface{}
		index    int
		value    interface{}
		expected interface{}
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, 1, 9, []int{1, 9, 2, 3}},
		{[]int{1, 2, 3}, 0, 9, []int{9, 1, 2, 3}},
		{[]int{1, 2, 3}, 3, 9, []int{1, 2, 3, 9}},
		{[]int{1, 2, 3}, -1, 9, []int{9, 1, 2, 3}},
		{[]int{1, 2, 3}, 4, 9, []int{1, 2, 3, 9}},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, 1, "x", []string{"a", "x", "b", "c"}},
		{[]string{"a", "b", "c"}, 0, "x", []string{"x", "a", "b", "c"}},
		{[]string{"a", "b", "c"}, 3, "x", []string{"a", "b", "c", "x"}},
		{[]string{"a", "b", "c"}, -1, "x", []string{"x", "a", "b", "c"}},
		{[]string{"a", "b", "c"}, 4, "x", []string{"a", "b", "c", "x"}},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result := arrayhelper.InsertSafe(arr, tt.index, tt.value.(int))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertSafe(%v, %d, %v) = %v; want %v", arr, tt.index, tt.value, result, tt.expected)
			}
		case []string:
			result := arrayhelper.InsertSafe(arr, tt.index, tt.value.(string))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertSafe(%v, %d, %v) = %v; want %v", arr, tt.index, tt.value, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}

func TestInsertWithErr(t *testing.T) {
	tests := []struct {
		arr      interface{}
		index    int
		value    interface{}
		expected interface{}
		err      error
	}{
		// 测试用例 int 类型
		{[]int{1, 2, 3}, 1, 9, []int{1, 9, 2, 3}, nil},
		{[]int{1, 2, 3}, 0, 9, []int{9, 1, 2, 3}, nil},
		{[]int{1, 2, 3}, 3, 9, []int{1, 2, 3, 9}, nil},
		{[]int{1, 2, 3}, -1, 9, []int{1, 2, 3}, arrayhelper.ErrOutOfRange},
		{[]int{1, 2, 3}, 4, 9, []int{1, 2, 3}, arrayhelper.ErrOutOfRange},

		// 测试用例 string 类型
		{[]string{"a", "b", "c"}, 1, "x", []string{"a", "x", "b", "c"}, nil},
		{[]string{"a", "b", "c"}, 0, "x", []string{"x", "a", "b", "c"}, nil},
		{[]string{"a", "b", "c"}, 3, "x", []string{"a", "b", "c", "x"}, nil},
		{[]string{"a", "b", "c"}, -1, "x", []string{"a", "b", "c"}, arrayhelper.ErrOutOfRange},
		{[]string{"a", "b", "c"}, 4, "x", []string{"a", "b", "c"}, arrayhelper.ErrOutOfRange},
	}

	for _, tt := range tests {
		switch arr := tt.arr.(type) {
		case []int:
			result, err := arrayhelper.InsertWithErr(arr, tt.index, tt.value.(int))
			if !errors.Is(err, tt.err) {
				t.Errorf("InsertWithErr(%v, %d, %v) error = %v; want %v", arr, tt.index, tt.value, err, tt.err)
			}
			if err == nil && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertWithErr(%v, %d, %v) = %v; want %v", arr, tt.index, tt.value, result, tt.expected)
			}
		case []string:
			result, err := arrayhelper.InsertWithErr(arr, tt.index, tt.value.(string))
			if !errors.Is(err, tt.err) {
				t.Errorf("InsertWithErr(%v, %d, %v) error = %v; want %v", arr, tt.index, tt.value, err, tt.err)
			}
			if err == nil && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertWithErr(%v, %d, %v) = %v; want %v", arr, tt.index, tt.value, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.arr)
		}
	}
}
