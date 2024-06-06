package mathhelper_test

import (
	"testing"

	mathhelper "github.com/jbterrylin/go-helper/mathHelper"
)

func TestMax(t *testing.T) {
	type TestCase[T mathhelper.Comparable] struct {
		Name     string
		A, B     T
		Expected T
	}

	intCases := []TestCase[int]{
		{"Int: a > b", 3, 2, 3},
		{"Int: a < b", 1, 2, 2},
		{"Int: a == b", 2, 2, 2},
	}

	floatCases := []TestCase[float64]{
		{"Float: a > b", 3.5, 2.5, 3.5},
		{"Float: a < b", 1.5, 2.5, 2.5},
		{"Float: a == b", 2.5, 2.5, 2.5},
	}

	stringCases := []TestCase[string]{
		{"String: a > b", "b", "a", "b"},
		{"String: a < b", "a", "b", "b"},
		{"String: a == b", "a", "a", "a"},
	}

	for _, testCase := range intCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := mathhelper.Max(testCase.A, testCase.B)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}

	for _, testCase := range floatCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := mathhelper.Max(testCase.A, testCase.B)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}

	for _, testCase := range stringCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := mathhelper.Max(testCase.A, testCase.B)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
