package maphelper_test

import (
	"testing"

	maphelper "github.com/jbterrylin/go-helper/mapHelper"
)

func TestKeyEqual(t *testing.T) {
	testCases := []struct {
		Name      string
		FirstMap  map[string]interface{}
		SecondMap map[string]interface{}
		Expected  bool
	}{
		{
			Name: "Same keys",
			FirstMap: map[string]interface{}{
				"key1": nil,
				"key2": nil,
				"key3": nil,
			},
			SecondMap: map[string]interface{}{
				"key1": nil,
				"key2": nil,
				"key3": nil,
			},
			Expected: true,
		},
		{
			Name: "Different keys",
			FirstMap: map[string]interface{}{
				"key1": nil,
				"key2": nil,
			},
			SecondMap: map[string]interface{}{
				"key1": nil,
				"key3": nil,
			},
			Expected: false,
		},
		{
			Name: "Different lengths",
			FirstMap: map[string]interface{}{
				"key1": nil,
				"key2": nil,
				"key3": nil,
			},
			SecondMap: map[string]interface{}{
				"key1": nil,
				"key2": nil,
			},
			Expected: false,
		},
		{
			Name:      "Empty maps",
			FirstMap:  map[string]interface{}{},
			SecondMap: map[string]interface{}{},
			Expected:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := maphelper.KeyEqual(testCase.FirstMap, testCase.SecondMap)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
