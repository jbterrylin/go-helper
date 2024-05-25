package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
	pointerhelper "github.com/jbterrylin/go-helper/pointerHelper"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		firstArr  interface{}
		secondArr interface{}
		expected  bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}, true},
		{[]float64{1.1, 2.2, 3.3}, []float64{3.3, 2.2, 1.1}, false},
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2}, false},
	}

	for _, tt := range tests {
		switch firstArr := tt.firstArr.(type) {
		case []int:
			secondArr := tt.secondArr.([]int)
			result := arrayhelper.Equal(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("Equal(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		case []string:
			secondArr := tt.secondArr.([]string)
			result := arrayhelper.Equal(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("Equal(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		case []float64:
			secondArr := tt.secondArr.([]float64)
			result := arrayhelper.Equal(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("Equal(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.firstArr)
		}
	}
}

func TestEqualRef(t *testing.T) {
	tests := []struct {
		firstArr  interface{}
		secondArr interface{}
		expected  bool
	}{
		{[]*int{pointerhelper.Pointer(1)}, []*int{pointerhelper.Pointer(1)}, true},
		{[]*int{pointerhelper.Pointer(1)}, []*int{pointerhelper.Pointer(2)}, false},
		{[]*int{pointerhelper.Pointer(1), pointerhelper.Pointer(2)}, []*int{pointerhelper.Pointer(1), pointerhelper.Pointer(2)}, true},
		{[]*int{pointerhelper.Pointer(1), nil}, []*int{pointerhelper.Pointer(1), nil}, true},
		{[]*int{nil, pointerhelper.Pointer(2)}, []*int{nil, pointerhelper.Pointer(2)}, true},
		{[]*int{nil, pointerhelper.Pointer(2)}, []*int{pointerhelper.Pointer(1), pointerhelper.Pointer(2)}, false},
		{[]*int{pointerhelper.Pointer(1), pointerhelper.Pointer(2)}, []*int{nil, pointerhelper.Pointer(2)}, false},
		{[]*int{pointerhelper.Pointer(1)}, []*int{pointerhelper.Pointer(1), pointerhelper.Pointer(2)}, false},

		{[]*string{pointerhelper.Pointer("a")}, []*string{pointerhelper.Pointer("a")}, true},
		{[]*string{pointerhelper.Pointer("a")}, []*string{pointerhelper.Pointer("b")}, false},
		{[]*string{pointerhelper.Pointer("a"), pointerhelper.Pointer("b")}, []*string{pointerhelper.Pointer("a"), pointerhelper.Pointer("b")}, true},
		{[]*string{pointerhelper.Pointer("a"), nil}, []*string{pointerhelper.Pointer("a"), nil}, true},
		{[]*string{nil, pointerhelper.Pointer("b")}, []*string{nil, pointerhelper.Pointer("b")}, true},
		{[]*string{nil, pointerhelper.Pointer("b")}, []*string{pointerhelper.Pointer("a"), pointerhelper.Pointer("b")}, false},
		{[]*string{pointerhelper.Pointer("a"), pointerhelper.Pointer("b")}, []*string{nil, pointerhelper.Pointer("b")}, false},
		{[]*string{pointerhelper.Pointer("a")}, []*string{pointerhelper.Pointer("a"), pointerhelper.Pointer("b")}, false},

		{[]*float64{pointerhelper.Pointer(1.1)}, []*float64{pointerhelper.Pointer(1.1)}, true},
		{[]*float64{pointerhelper.Pointer(1.1)}, []*float64{pointerhelper.Pointer(2.2)}, false},
		{[]*float64{pointerhelper.Pointer(1.1), pointerhelper.Pointer(2.2)}, []*float64{pointerhelper.Pointer(1.1), pointerhelper.Pointer(2.2)}, true},
		{[]*float64{pointerhelper.Pointer(1.1), nil}, []*float64{pointerhelper.Pointer(1.1), nil}, true},
		{[]*float64{nil, pointerhelper.Pointer(2.2)}, []*float64{nil, pointerhelper.Pointer(2.2)}, true},
		{[]*float64{nil, pointerhelper.Pointer(2.2)}, []*float64{pointerhelper.Pointer(1.1), pointerhelper.Pointer(2.2)}, false},
		{[]*float64{pointerhelper.Pointer(1.1), pointerhelper.Pointer(2.2)}, []*float64{nil, pointerhelper.Pointer(2.2)}, false},
		{[]*float64{pointerhelper.Pointer(1.1)}, []*float64{pointerhelper.Pointer(1.1), pointerhelper.Pointer(2.2)}, false},
	}

	for _, tt := range tests {
		switch firstArr := tt.firstArr.(type) {
		case []*int:
			secondArr := tt.secondArr.([]*int)
			result := arrayhelper.EqualRef(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("EqualRef(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		case []*string:
			secondArr := tt.secondArr.([]*string)
			result := arrayhelper.EqualRef(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("EqualRef(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		case []*float64:
			secondArr := tt.secondArr.([]*float64)
			result := arrayhelper.EqualRef(firstArr, secondArr)
			if result != tt.expected {
				t.Errorf("EqualRef(%v, %v) = %v; want %v", firstArr, secondArr, result, tt.expected)
			}
		default:
			t.Fatalf("unsupported type: %T", tt.firstArr)
		}
	}
}
