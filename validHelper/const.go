package validhelper

import "errors"

const NOT_EMPTY = "notEmpty"
const REGEXP = "regexp"
const LT = "LT"
const LE = "LE"
const EQ = "EQ"
const NE = "NE"
const GE = "GE"
const GT = "GT"

const CUSTOM_RULE = "CUSTOM_RULE:"

var compareMap = map[string]bool{
	LT: true,
	LE: true,
	EQ: true,
	NE: true,
	GE: true,
	GT: true,
}

var ErrIsEmpty = errors.New("is empty")
var ErrRegexp = errors.New("regexp fail")
var ErrLenNotValid = errors.New("len is not valid")
