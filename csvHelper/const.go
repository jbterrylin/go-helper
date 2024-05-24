package csvhelper

import "errors"

const CsvTag = "l_csv"
const CsvHeader = "header:"
const CsvSort = "sort:"
const CsvHide = "hide:"
const CsvDefaultConverter = "coverter:"

var ErrNotArray = errors.New("not values")
var ErrEmptyArray = errors.New("empty values")
var ErrValNotStruct = errors.New("value not struct")
var ErrValsNotSlicePointer = errors.New("values must be a pointer to a slice")
