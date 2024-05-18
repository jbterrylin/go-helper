package filehelper

import (
	"errors"
)

var ErrPathIsPointToFile = errors.New("path is point to file")
var ErrPathIsPointToDir = errors.New("path is point to directory")
