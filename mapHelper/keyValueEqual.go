package maphelper

import "reflect"

func KeyValueEqual[T comparable](firstMap, secondMap map[T]interface{}) bool {
	if len(firstMap) != len(secondMap) {
		return false
	}

	for key, firstValue := range firstMap {
		secondValue, exists := secondMap[key]
		if !exists || !reflect.DeepEqual(firstValue, secondValue) {
			return false
		}
	}
	return true
}
