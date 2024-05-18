package maphelper

import (
	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
)

func KeyEqual[T comparable](firstMap, secondMap map[T]interface{}) bool {
	if len(firstMap) != len(secondMap) {
		return false
	}

	firstKeys := Keys(firstMap)
	secondKeys := Keys(secondMap)

	for _, val := range firstKeys {
		index := arrayhelper.IndexOf(secondKeys, val)
		if index != -1 {
			arrayhelper.Splice(&secondKeys, index, 1)
		} else {
			return false
		}
	}
	return true
}
