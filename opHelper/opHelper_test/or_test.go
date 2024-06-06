package ophelper_test

import (
	"testing"

	ophelper "github.com/jbterrylin/go-helper/opHelper"
)

func TestOr(t *testing.T) {
	type TestCase[T any] struct {
		Name      string
		Value     T
		ElseValue T
		Condition func(T) bool
		Expected  T
	}

	intCases := []TestCase[int]{
		{"Int: Non-zero value", 1, 2, nil, 1},
		{"Int: Zero value", 0, 2, nil, 2},
	}

	floatCases := []TestCase[float64]{
		{"Float: Non-zero value", 1.1, 2.2, nil, 1.1},
		{"Float: Zero value", 0.0, 2.2, nil, 2.2},
	}

	stringCases := []TestCase[string]{
		{"String: Non-empty value", "a", "b", nil, "a"},
		{"String: Empty value", "", "b", nil, "b"},
	}

	for _, testCase := range intCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ophelper.Or(testCase.Value, testCase.ElseValue)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}

	for _, testCase := range floatCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ophelper.Or(testCase.Value, testCase.ElseValue)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}

	for _, testCase := range stringCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ophelper.Or(testCase.Value, testCase.ElseValue)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}

func TestOrByCond(t *testing.T) {
	type TestCase[T any] struct {
		Name      string
		Value     T
		ElseValue T
		Condition func(T) bool
		Expected  T
	}

	cond := func(value int) bool { return value < 0 }

	intCases := []TestCase[int]{
		{"Int: Condition false", 1, 2, cond, 1},
		{"Int: Condition true", -1, 2, cond, 2},
	}

	condFloat := func(value float64) bool { return value < 0 }

	floatCases := []TestCase[float64]{
		{"Float: Condition false", 1.1, 2.2, condFloat, 1.1},
		{"Float: Condition true", -1.1, 2.2, condFloat, 2.2},
	}

	condString := func(value string) bool { return value == "" }

	stringCases := []TestCase[string]{
		{"String: Condition false", "a", "b", condString, "a"},
		{"String: Condition true", "", "b", condString, "b"},
	}

	for _, testCase := range intCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ophelper.OrByCond(testCase.Value, testCase.ElseValue, testCase.Condition)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}

	for _, testCase := range floatCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ophelper.OrByCond(testCase.Value, testCase.ElseValue, testCase.Condition)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}

	for _, testCase := range stringCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ophelper.OrByCond(testCase.Value, testCase.ElseValue, testCase.Condition)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
