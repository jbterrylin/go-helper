package pointerhelper_test

import (
	"testing"

	pointerhelper "github.com/jbterrylin/go-helper/pointerHelper"
)

func TestPointer(t *testing.T) {
	type TestCase[T any] struct {
		Name     string
		Input    T
		Expected T
	}

	intCases := []TestCase[int]{
		{"Int: Positive", 42, 42},
		{"Int: Zero", 0, 0},
		{"Int: Negative", -42, -42},
	}

	floatCases := []TestCase[float64]{
		{"Float: Positive", 3.14, 3.14},
		{"Float: Zero", 0.0, 0.0},
		{"Float: Negative", -3.14, -3.14},
	}

	stringCases := []TestCase[string]{
		{"String: Non-empty", "hello", "hello"},
		{"String: Empty", "", ""},
	}

	for _, testCase := range intCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := pointerhelper.Pointer(testCase.Input)
			if *result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, *result)
			}
		})
	}

	for _, testCase := range floatCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := pointerhelper.Pointer(testCase.Input)
			if *result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, *result)
			}
		})
	}

	for _, testCase := range stringCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := pointerhelper.Pointer(testCase.Input)
			if *result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, *result)
			}
		})
	}
}
