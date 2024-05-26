package formathelper_test

import (
	"testing"

	formathelper "github.com/jbterrylin/go-helper/formatHelper"
)

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		value          interface{}
		decimalPlaces  int
		expectedResult interface{}
	}{
		{123.456, 2, 123.46},
		{123.454, 2, 123.45},
		{123, 2, 123},
		{123, 0, 123},
		{uint(123), 2, uint(123)},
		{-123.456, 2, -123.46},
		{-123.454, 2, -123.45},
	}

	for _, test := range tests {
		var result interface{}
		switch v := test.value.(type) {
		case int:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case int8:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case int16:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case int32:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case int64:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case uint:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case uint8:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case uint16:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case uint32:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case uint64:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case float32:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		case float64:
			result = formathelper.RoundFloat(v, test.decimalPlaces)
		default:
			t.Errorf("Unsupported type: %T", test.value)
		}

		if result != test.expectedResult {
			t.Errorf("RoundFloat(%v, %d) = %v; want %v", test.value, test.decimalPlaces, result, test.expectedResult)
		}
	}
}

func TestCeilFloat(t *testing.T) {
	tests := []struct {
		value          interface{}
		decimalPlaces  int
		expectedResult interface{}
	}{
		{123.456, 2, 123.46},
		{123.451, 2, 123.46},
		{123, 2, 123},
		{123, 0, 123},
		{uint(123), 2, uint(123)},
		{-123.456, 2, -123.45},
		{-123.451, 2, -123.45},
	}

	for _, test := range tests {
		var result interface{}
		switch v := test.value.(type) {
		case int:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case int8:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case int16:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case int32:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case int64:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case uint:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case uint8:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case uint16:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case uint32:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case uint64:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case float32:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		case float64:
			result = formathelper.CeilFloat(v, test.decimalPlaces)
		default:
			t.Errorf("Unsupported type: %T", test.value)
		}

		if result != test.expectedResult {
			t.Errorf("CeilFloat(%v, %d) = %v; want %v", test.value, test.decimalPlaces, result, test.expectedResult)
		}
	}
}

func TestFloorFloat(t *testing.T) {
	tests := []struct {
		value          interface{}
		decimalPlaces  int
		expectedResult interface{}
	}{
		{123.456, 2, 123.45},
		{123.459, 2, 123.45},
		{123, 2, 123},
		{123, 0, 123},
		{uint(123), 2, uint(123)},
		{-123.456, 2, -123.46},
		{-123.459, 2, -123.46},
	}

	for _, test := range tests {
		var result interface{}
		switch v := test.value.(type) {
		case int:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case int8:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case int16:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case int32:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case int64:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case uint:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case uint8:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case uint16:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case uint32:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case uint64:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case float32:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		case float64:
			result = formathelper.FloorFloat(v, test.decimalPlaces)
		default:
			t.Errorf("Unsupported type: %T", test.value)
		}

		if result != test.expectedResult {
			t.Errorf("FloorFloat(%v, %d) = %v; want %v", test.value, test.decimalPlaces, result, test.expectedResult)
		}
	}
}

func TestTruncateFloat(t *testing.T) {
	tests := []struct {
		value          interface{}
		decimalPlaces  int
		expectedResult interface{}
	}{
		{123.456, 2, 123.45},
		{123.459, 2, 123.45},
		{123, 2, 123},
		{123, 0, 123},
		{uint(123), 2, uint(123)},
		{-123.456, 2, -123.45},
		{-123.459, 2, -123.45},
	}

	for _, test := range tests {
		var result interface{}
		switch v := test.value.(type) {
		case int:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case int8:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case int16:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case int32:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case int64:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case uint:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case uint8:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case uint16:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case uint32:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case uint64:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case float32:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		case float64:
			result = formathelper.TruncateFloat(v, test.decimalPlaces)
		default:
			t.Errorf("Unsupported type: %T", test.value)
		}

		if result != test.expectedResult {
			t.Errorf("TruncateFloat(%v, %d) = %v; want %v", test.value, test.decimalPlaces, result, test.expectedResult)
		}
	}
}
