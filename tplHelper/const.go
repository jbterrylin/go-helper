package tplhelper

import "errors"

const defaultTempDir = "tplTmp"
const fileNamePlaceHolder = "${ReplaceFileName}"

var ErrFileExist = errors.New("file exist")
