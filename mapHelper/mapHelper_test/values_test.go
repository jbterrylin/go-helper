package maphelper_test

import (
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
	maphelper "github.com/jbterrylin/go-helper/mapHelper"
)

func TestValues(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    map[string]int
		Expected []int
	}{
		{
			Name:     "Empty map",
			Input:    map[string]int{},
			Expected: []int{},
		},
		{
			Name: "Single value",
			Input: map[string]int{
				"key1": 1,
			},
			Expected: []int{1},
		},
		{
			Name: "Multiple values",
			Input: map[string]int{
				"key1": 1,
				"key2": 2,
				"key3": 3,
			},
			Expected: []int{1, 2, 3},
		},
		{
			Name: "Non-sequential values",
			Input: map[string]int{
				"key1": 10,
				"key2": 20,
				"key3": 30,
			},
			Expected: []int{10, 20, 30},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := maphelper.Values(testCase.Input)
			if !arrayhelper.Equal(result, testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
