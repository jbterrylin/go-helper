package csvhelper_test

import (
	"reflect"
	"testing"

	csvhelper "github.com/jbterrylin/go-helper/csvHelper"
)

func TestCsvToStructs(t *testing.T) {
	// Define a struct for testing
	type TestStruct struct {
		Name   string  `l_csv:"header:Name"`
		Age    int     `l_csv:"header:Age"`
		Salary float64 `l_csv:"header:Salary1"`
		Active bool
	}

	csvData := [][]string{
		{"Name", "Age", "Salary1", "Active"},
		{"Alice", "30", "50000.50", "TRUE"},
		{"Bob", "25", "40000.00", "FALSE"},
	}

	var result []TestStruct
	err := csvhelper.CsvToStructs(csvData, reflect.ValueOf(&result))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []TestStruct{
		{Name: "Alice", Age: 30, Salary: 50000.50, Active: true},
		{Name: "Bob", Age: 25, Salary: 40000.00, Active: false},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
