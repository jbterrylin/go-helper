package arrayhelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func testSpliceSlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestSplice(t *testing.T) {
	tests := []struct {
		slice       []float64
		start       int
		deleteCount int
		items       []float64
		expected    []float64
		remaining   []float64
	}{
		// 测试用例 float64 类型
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 2, 2, []float64{8.8, 9.9}, []float64{3.3, 4.4}, []float64{1.1, 2.2, 8.8, 9.9, 5.5}},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, -2, 1, []float64{8.8, 9.9}, []float64{4.4}, []float64{1.1, 2.2, 3.3, 8.8, 9.9, 5.5}},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 1, 3, []float64{}, []float64{2.2, 3.3, 4.4}, []float64{1.1, 5.5}},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 1, -1, []float64{8.8, 9.9}, []float64{}, []float64{1.1, 8.8, 9.9, 2.2, 3.3, 4.4, 5.5}},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 5, 1, []float64{8.8, 9.9}, []float64{}, []float64{1.1, 2.2, 3.3, 4.4, 5.5, 8.8, 9.9}},
	}

	for _, tt := range tests {
		sliceCopy := make([]float64, len(tt.slice))
		copy(sliceCopy, tt.slice)
		removed := arrayhelper.Splice(&sliceCopy, tt.start, tt.deleteCount, tt.items...)
		if !testSpliceSlicesEqual(removed, tt.expected) {
			t.Errorf("Splice(%v, %d, %d, %v) = %v; want %v", tt.slice, tt.start, tt.deleteCount, tt.items, removed, tt.expected)
		}
		if !testSpliceSlicesEqual(sliceCopy, tt.remaining) {
			t.Errorf("Splice remaining = %v; want %v", sliceCopy, tt.remaining)
		}
	}
}
