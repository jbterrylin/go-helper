package maphelper

import (
	"bytes"
	"encoding/gob"
)

func CopyMap(src, desc interface{}) {
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(src)
	gob.NewDecoder(buf).Decode(desc)
}
