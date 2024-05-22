package gormhelper

import (
	"strings"

	"gorm.io/gorm"
)

func joinAll(db *gorm.DB, joinMap, preloadMap map[string][]interface{}) *gorm.DB {
	for key, val := range joinMap {
		db = db.Joins(key, val...)
	}
	for key, val := range preloadMap {
		db = db.Preload(key, val...)
	}
	return db
}

func joinByJoinStrings(db *gorm.DB, joinMap, preloadMap map[string][]interface{}, joinStrings ...string) *gorm.DB {
	for _, joinString := range joinStrings {
		if val, ok := joinMap[joinString]; ok {
			db = db.Joins(joinString, val...)
		} else {
			val := preloadMap[joinString]
			db = db.Preload(joinString, val...)
		}
	}
	return db
}

func except(joinMap, preloadMap map[string][]interface{}, joinStrings ...string) {
	for _, joinString := range joinStrings {
		if strings.HasPrefix(joinString, EXCEPT) {
			key := strings.TrimPrefix(joinString, EXCEPT)
			delete(joinMap, key)
			delete(preloadMap, key)
		}
	}
}

func Join(db *gorm.DB, joinMap, preloadMap map[string][]interface{}, joinStrings ...string) *gorm.DB {
	var specialHandle bool
	if len(joinStrings) == 0 {
		db = joinAll(db, joinMap, preloadMap)
		return db
	}
	if len(joinStrings) == 1 {
		switch joinStrings[0] {
		case NO_JOIN:
			joinMap = map[string][]interface{}{}
			preloadMap = map[string][]interface{}{}
			specialHandle = true
		case ONLY_JOIN:
			preloadMap = map[string][]interface{}{}
			specialHandle = true
		case ONLY_PRELOAD:
			joinMap = map[string][]interface{}{}
			specialHandle = true
		}
	}
	if specialHandle {
		db = joinAll(db, joinMap, preloadMap)
		return db
	}
	if strings.HasPrefix(joinStrings[0], EXCEPT) {
		except(joinMap, preloadMap, joinStrings...)
		db = joinAll(db, joinMap, preloadMap)
		return db
	}

	db = joinByJoinStrings(db, joinMap, preloadMap, joinStrings...)

	return db
}
