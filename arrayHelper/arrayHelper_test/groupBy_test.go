package arrayhelper_test

import (
	"reflect"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	tests := []struct {
		name     string
		input    interface{}
		getKey   interface{}
		expected interface{}
	}{
		{
			name: "Group by age",
			input: []Person{
				{Name: "Alice", Age: 30},
				{Name: "Bob", Age: 40},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 40},
				{Name: "Eve", Age: 50},
			},
			getKey: func(p Person) int { return p.Age },
			expected: map[int][]Person{
				30: {
					{Name: "Alice", Age: 30},
					{Name: "Charlie", Age: 30},
				},
				40: {
					{Name: "Bob", Age: 40},
					{Name: "David", Age: 40},
				},
				50: {
					{Name: "Eve", Age: 50},
				},
			},
		},
		{
			name:   "Group by first letter",
			input:  []string{"apple", "banana", "apricot", "blueberry", "avocado"},
			getKey: func(s string) string { return string(s[0]) },
			expected: map[string][]string{
				"a": {"apple", "apricot", "avocado"},
				"b": {"banana", "blueberry"},
			},
		},
		{
			name:     "Group empty slice",
			input:    []int{},
			getKey:   func(i int) int { return i },
			expected: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch input := tt.input.(type) {
			case []Person:
				getKey := tt.getKey.(func(Person) int)
				result := arrayhelper.GroupBy(input, getKey)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			case []string:
				getKey := tt.getKey.(func(string) string)
				result := arrayhelper.GroupBy(input, getKey)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			case []int:
				getKey := tt.getKey.(func(int) int)
				result := arrayhelper.GroupBy(input, getKey)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			default:
				t.Fatalf("unsupported input type: %T", tt.input)
			}
		})
	}
}

func TestGroupByAndReshape(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	tests := []struct {
		name     string
		input    interface{}
		getKey   interface{}
		reshape  interface{}
		expected interface{}
	}{
		{
			name: "Group by age and reshape to names",
			input: []Person{
				{Name: "Alice", Age: 30},
				{Name: "Bob", Age: 40},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 40},
				{Name: "Eve", Age: 50},
			},
			getKey:  func(p Person) int { return p.Age },
			reshape: func(p Person) string { return p.Name },
			expected: map[int][]string{
				30: {"Alice", "Charlie"},
				40: {"Bob", "David"},
				50: {"Eve"},
			},
		},
		{
			name:    "Group by first letter and reshape to lengths",
			input:   []string{"apple", "banana", "apricot", "blueberry", "avocado"},
			getKey:  func(s string) string { return string(s[0]) },
			reshape: func(s string) int { return len(s) },
			expected: map[string][]int{
				"a": {5, 7, 7},
				"b": {6, 9},
			},
		},
		{
			name:     "Group empty slice and reshape to same",
			input:    []int{},
			getKey:   func(i int) int { return i },
			reshape:  func(i int) int { return i },
			expected: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch input := tt.input.(type) {
			case []Person:
				getKey := tt.getKey.(func(Person) int)
				reshape := tt.reshape.(func(Person) string)
				result := arrayhelper.GroupByAndReshape(input, getKey, reshape)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			case []string:
				getKey := tt.getKey.(func(string) string)
				reshape := tt.reshape.(func(string) int)
				result := arrayhelper.GroupByAndReshape(input, getKey, reshape)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			case []int:
				getKey := tt.getKey.(func(int) int)
				reshape := tt.reshape.(func(int) int)
				result := arrayhelper.GroupByAndReshape(input, getKey, reshape)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			default:
				t.Fatalf("unsupported input type: %T", tt.input)
			}
		})
	}
}
