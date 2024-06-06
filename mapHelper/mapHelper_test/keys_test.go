package maphelper_test

import (
	"sort"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
	maphelper "github.com/jbterrylin/go-helper/mapHelper"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    map[string]int
		Expected []string
	}{
		{
			Name:     "Empty map",
			Input:    map[string]int{},
			Expected: []string{},
		},
		{
			Name: "Single key",
			Input: map[string]int{
				"key1": 1,
			},
			Expected: []string{"key1"},
		},
		{
			Name: "Multiple keys",
			Input: map[string]int{
				"key1": 1,
				"key2": 2,
				"key3": 3,
			},
			Expected: []string{"key1", "key2", "key3"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := maphelper.Keys(testCase.Input)
			sort.Strings(result)
			sort.Strings(testCase.Expected)
			if !arrayhelper.Equal(result, testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
