package maphelper_test

import (
	"reflect"
	"testing"

	maphelper "github.com/jbterrylin/go-helper/mapHelper"
)

func TestCopyMap(t *testing.T) {
	type SimpleStruct struct {
		ID    int
		Value string
	}

	type ComplexStruct struct {
		ID      int
		Name    string
		Details map[string]string
	}

	testCases := []struct {
		Name     string
		Src      interface{}
		Dest     interface{}
		Expected interface{}
	}{
		{
			Name: "Simple map copy",
			Src: map[string]SimpleStruct{
				"key1": {ID: 1, Value: "A"},
				"key2": {ID: 2, Value: "B"},
				"key3": {ID: 3, Value: "C"},
			},
			Dest: &map[string]SimpleStruct{},
			Expected: map[string]SimpleStruct{
				"key1": {ID: 1, Value: "A"},
				"key2": {ID: 2, Value: "B"},
				"key3": {ID: 3, Value: "C"},
			},
		},
		{
			Name: "Complex map copy",
			Src: map[string]ComplexStruct{
				"key1": {ID: 1, Name: "Item1", Details: map[string]string{"Detail1": "A", "Detail2": "B"}},
				"key2": {ID: 2, Name: "Item2", Details: map[string]string{"Detail3": "C", "Detail4": "D"}},
			},
			Dest: &map[string]ComplexStruct{},
			Expected: map[string]ComplexStruct{
				"key1": {ID: 1, Name: "Item1", Details: map[string]string{"Detail1": "A", "Detail2": "B"}},
				"key2": {ID: 2, Name: "Item2", Details: map[string]string{"Detail3": "C", "Detail4": "D"}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			maphelper.CopyMap(testCase.Src, testCase.Dest)
			if !reflect.DeepEqual(reflect.ValueOf(testCase.Dest).Elem().Interface(), testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, reflect.ValueOf(testCase.Dest).Elem().Interface())
			}
		})
	}
}
