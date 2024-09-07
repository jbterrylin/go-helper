package gormhelper_test

import (
	"testing"

	gormhelper "github.com/jbterrylin/go-helper/gormHelper"
)

type MainStruct struct {
	ID            int
	DerivedStruct DerivedStruct
}

type DerivedStruct struct {
	MainStructID int
	Value        string
}

func TestCrossJoin(t *testing.T) {
	testCases := []struct {
		Name           string
		MainStructs    []MainStruct
		DerivedStructs []DerivedStruct
		Expected       []MainStruct
	}{
		{
			Name: "Simple match",
			MainStructs: []MainStruct{
				{ID: 1}, {ID: 2}, {ID: 3},
			},
			DerivedStructs: []DerivedStruct{
				{MainStructID: 1, Value: "A"},
				{MainStructID: 2, Value: "B"},
				{MainStructID: 3, Value: "C"},
			},
			Expected: []MainStruct{
				{ID: 1, DerivedStruct: DerivedStruct{MainStructID: 1, Value: "A"}},
				{ID: 2, DerivedStruct: DerivedStruct{MainStructID: 2, Value: "B"}},
				{ID: 3, DerivedStruct: DerivedStruct{MainStructID: 3, Value: "C"}},
			},
		},
		{
			Name: "No derived structs",
			MainStructs: []MainStruct{
				{ID: 4}, {ID: 5},
			},
			DerivedStructs: []DerivedStruct{},
			Expected: []MainStruct{
				{ID: 4}, {ID: 5},
			},
		},
		{
			Name: "Partial match",
			MainStructs: []MainStruct{
				{ID: 1}, {ID: 4}, {ID: 3},
			},
			DerivedStructs: []DerivedStruct{
				{MainStructID: 1, Value: "A"},
				{MainStructID: 2, Value: "B"},
				{MainStructID: 3, Value: "C"},
			},
			Expected: []MainStruct{
				{ID: 1, DerivedStruct: DerivedStruct{MainStructID: 1, Value: "A"}},
				{ID: 4},
				{ID: 3, DerivedStruct: DerivedStruct{MainStructID: 3, Value: "C"}},
			},
		},
		{
			Name: "Multiple derived structs with same ID",
			MainStructs: []MainStruct{
				{ID: 1}, {ID: 2},
			},
			DerivedStructs: []DerivedStruct{
				{MainStructID: 1, Value: "A1"},
				{MainStructID: 1, Value: "A2"},
				{MainStructID: 2, Value: "B"},
			},
			Expected: []MainStruct{
				{ID: 1, DerivedStruct: DerivedStruct{MainStructID: 1, Value: "A1"}},
				{ID: 2, DerivedStruct: DerivedStruct{MainStructID: 2, Value: "B"}},
			},
		},
		{
			Name: "Derived structs out of order",
			MainStructs: []MainStruct{
				{ID: 1}, {ID: 2}, {ID: 3},
			},
			DerivedStructs: []DerivedStruct{
				{MainStructID: 2, Value: "B"},
				{MainStructID: 1, Value: "A"},
				{MainStructID: 3, Value: "C"},
			},
			Expected: []MainStruct{
				{ID: 1, DerivedStruct: DerivedStruct{MainStructID: 1, Value: "A"}},
				{ID: 2, DerivedStruct: DerivedStruct{MainStructID: 2, Value: "B"}},
				{ID: 3, DerivedStruct: DerivedStruct{MainStructID: 3, Value: "C"}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			getDerivedId := func(mainStruct MainStruct) interface{} {
				return mainStruct.ID
			}

			getDerivedStructs := func(ids []interface{}) ([]DerivedStruct, error) {
				return testCase.DerivedStructs, nil
			}

			setDerived := func(mainStruct *MainStruct, derivedStruct DerivedStruct) bool {
				if mainStruct.ID == derivedStruct.MainStructID {
					mainStruct.DerivedStruct = derivedStruct
					return true
				}
				return false
			}

			result, err := gormhelper.CrossJoin(testCase.MainStructs, getDerivedId, getDerivedStructs, setDerived)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			for i, mainStruct := range result {
				if mainStruct.DerivedStruct != testCase.Expected[i].DerivedStruct {
					t.Errorf("Expected DerivedStruct %+v, got %+v", testCase.Expected[i].DerivedStruct, mainStruct.DerivedStruct)
				}
			}
		})
	}
}
