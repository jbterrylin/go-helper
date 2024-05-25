package arrayhelper_test

import (
	"errors"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestPop(t *testing.T) {
	tests := []struct {
		slice    interface{}
		expected interface{}
		err      error
	}{
		// 测试用例 int 类型
		{&[]int{1, 2, 3}, 3, nil},
		{&[]int{}, nil, arrayhelper.ErrOutOfRange},

		// 测试用例 string 类型
		{&[]string{"a", "b", "c"}, "c", nil},
		{&[]string{}, nil, arrayhelper.ErrOutOfRange},

		// 测试用例 float64 类型
		{&[]float64{1.1, 2.2, 3.3}, 3.3, nil},
		{&[]float64{}, nil, arrayhelper.ErrOutOfRange},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case *[]int:
			if tt.err != nil {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Pop did not panic for %v", slice)
					}
				}()
				arrayhelper.Pop(slice)
			} else {
				result := arrayhelper.Pop(slice)
				if result != tt.expected {
					t.Errorf("Pop(%v) = %v; want %v", slice, result, tt.expected)
				}
			}
		case *[]string:
			if tt.err != nil {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Pop did not panic for %v", slice)
					}
				}()
				arrayhelper.Pop(slice)
			} else {
				result := arrayhelper.Pop(slice)
				if result != tt.expected {
					t.Errorf("Pop(%v) = %v; want %v", slice, result, tt.expected)
				}
			}
		case *[]float64:
			if tt.err != nil {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Pop did not panic for %v", slice)
					}
				}()
				arrayhelper.Pop(slice)
			} else {
				result := arrayhelper.Pop(slice)
				if result != tt.expected {
					t.Errorf("Pop(%v) = %v; want %v", slice, result, tt.expected)
				}
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}

func TestPopWithErr(t *testing.T) {
	tests := []struct {
		slice    interface{}
		expected interface{}
		err      error
	}{
		// 测试用例 int 类型
		{&[]int{1, 2, 3}, 3, nil},
		{&[]int{}, *new(int), arrayhelper.ErrOutOfRange},

		// 测试用例 string 类型
		{&[]string{"a", "b", "c"}, "c", nil},
		{&[]string{}, *new(string), arrayhelper.ErrOutOfRange},

		// 测试用例 float64 类型
		{&[]float64{1.1, 2.2, 3.3}, 3.3, nil},
		{&[]float64{}, *new(float64), arrayhelper.ErrOutOfRange},
	}

	for _, tt := range tests {
		switch slice := tt.slice.(type) {
		case *[]int:
			result, err := arrayhelper.PopWithErr(slice)
			if !errors.Is(err, tt.err) {
				t.Errorf("PopWithErr(%v) error = %v; want %v", slice, err, tt.err)
			}
			if err == nil && result != tt.expected {
				t.Errorf("PopWithErr(%v) = %v; want %v", slice, result, tt.expected)
			}
		case *[]string:
			result, err := arrayhelper.PopWithErr(slice)
			if !errors.Is(err, tt.err) {
				t.Errorf("PopWithErr(%v) error = %v; want %v", slice, err, tt.err)
			}
			if err == nil && result != tt.expected {
				t.Errorf("PopWithErr(%v) = %v; want %v", slice, result, tt.expected)
			}
		case *[]float64:
			result, err := arrayhelper.PopWithErr(slice)
			if !errors.Is(err, tt.err) {
				t.Errorf("PopWithErr(%v) error = %v; want %v", slice, err, tt.err)
			}
			if err == nil && result != tt.expected {
				t.Errorf("PopWithErr(%v) = %v; want %v", slice, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.slice)
		}
	}
}
