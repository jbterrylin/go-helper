package formathelper_test

import (
	"testing"

	formathelper "github.com/jbterrylin/go-helper/formatHelper"
)

func TestFirstUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.FirstUpper(tt.input)
		if result != tt.expected {
			t.Errorf("FirstUpper(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"Hello World", "HelloWorld"},
		{"hello_world", "HelloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToCamelCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToCamelCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello-world"},
		{"Hello World", "hello-world"},
		{"hello_world", "hello-world"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToKebabCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToKebabCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello_world"},
		{"Hello World", "hello_world"},
		{"hello-world", "hello_world"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToSnakeCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToSnakeCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToLowerCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"hello_world", "helloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToLowerCamelCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToLowerCamelCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToUpperCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"Hello World", "HelloWorld"},
		{"hello_world", "HelloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToUpperCamelCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToUpperCamelCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"Hello World", "HelloWorld"},
		{"hello_world", "HelloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToPascalCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToPascalCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"hello_world", "Hello_world"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToTitleCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToTitleCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToSpaceCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello world"},
		{"Hello_World", "Hello World"},
		{"hello-world", "hello world"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToSpaceCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToSpaceCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello\tworld"},
		{"Hello_World", "Hello\tWorld"},
		{"hello-world", "hello\tworld"},
		{"", ""},
	}

	for _, tt := range tests {
		result := formathelper.ToTabCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToTabCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
