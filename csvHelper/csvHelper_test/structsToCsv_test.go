package csvhelper_test

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"

	csvhelper "github.com/jbterrylin/go-helper/csvHelper"
)

func TestStructsToCsv(t *testing.T) {
	// Define a struct for testing
	type TestStruct struct {
		Name   string  `l_csv:"header:Name"`
		Age    int     `l_csv:"header:Age"`
		Salary float64 `l_csv:"header:Salary"`
		Active bool    `l_csv:"header:Active"`
	}

	// Prepare test data
	data := []TestStruct{
		{Name: "Alice", Age: 30, Salary: 50000.50, Active: true},
		{Name: "Bob", Age: 25, Salary: 40000.00, Active: false},
	}

	// Create a buffer to write the CSV data
	file, err := os.Create("test_output.csv")
	if err != nil {
		t.Fatalf("Failed to create CSV file: %v", err)
	}
	defer os.Remove("test_output.csv")

	writer := csv.NewWriter(file)

	// Define CSV settings
	csvSetting := csvhelper.CsvSetting{
		RequireHeader: true,
		IgnoreFields:  []string{},
		Converters:    map[string]*csvhelper.CsvConverter{},
	}

	// Convert the data to a reflect.Value
	values := reflect.ValueOf(&data)

	// Call StructsToCsv function
	err = csvhelper.StructsToCsv(writer, values, csvSetting)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	writer.Flush()
	file.Close()

	// Verify the output file
	expectedOutput := "Name,Age,Salary,Active\nAlice,30,50000.5,true\nBob,25,40000,false\n"
	output, err := os.ReadFile("test_output.csv")
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}

	if string(output) != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, string(output))
	}
}
