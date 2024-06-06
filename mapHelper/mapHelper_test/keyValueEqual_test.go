package maphelper_test

import (
	"testing"

	maphelper "github.com/jbterrylin/go-helper/mapHelper"
)

func TestKeyValueEqual(t *testing.T) {
	testCases := []struct {
		Name      string
		FirstMap  map[string]interface{}
		SecondMap map[string]interface{}
		Expected  bool
	}{
		{
			Name: "Same keys and values",
			FirstMap: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			SecondMap: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			Expected: true,
		},
		{
			Name: "Different values",
			FirstMap: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			SecondMap: map[string]interface{}{
				"key1": "value1",
				"key2": "differentValue",
			},
			Expected: false,
		},
		{
			Name: "Different keys",
			FirstMap: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			SecondMap: map[string]interface{}{
				"key1": "value1",
				"key3": "value2",
			},
			Expected: false,
		},
		{
			Name: "Different lengths",
			FirstMap: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			SecondMap: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
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
			result := maphelper.KeyValueEqual(testCase.FirstMap, testCase.SecondMap)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
