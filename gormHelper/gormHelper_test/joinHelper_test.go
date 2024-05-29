package gormhelper_test

import (
	"strings"
	"testing"

	arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"
	gormhelper "github.com/jbterrylin/go-helper/gormHelper"
	maphelper "github.com/jbterrylin/go-helper/mapHelper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestJoin(t *testing.T) {
	joinMap := map[string][]interface{}{
		"UserJoin1": {},
		"UserJoin2": {},
	}
	preloadMap := map[string][]interface{}{
		"UserPreload1": {},
		"UserPreload2": {},
	}

	tests := []struct {
		name             string
		joinMap          map[string][]interface{}
		preloadMap       map[string][]interface{}
		joinStrings      []string
		expectedJoins    []string
		expectedPreloads []string
	}{
		{
			name:             strings.Join([]string{gormhelper.NO_JOIN}, ","),
			joinStrings:      []string{gormhelper.NO_JOIN},
			expectedJoins:    []string{},
			expectedPreloads: []string{},
		},
		{
			name:             strings.Join([]string{gormhelper.ONLY_JOIN}, ","),
			joinStrings:      []string{gormhelper.ONLY_JOIN},
			expectedJoins:    []string{"UserJoin1", "UserJoin2"},
			expectedPreloads: []string{},
		},
		{
			name:             strings.Join([]string{gormhelper.ONLY_PRELOAD}, ","),
			joinStrings:      []string{gormhelper.ONLY_PRELOAD},
			expectedJoins:    []string{},
			expectedPreloads: []string{"UserPreload1", "UserPreload2"},
		},
		{
			name:             strings.Join([]string{gormhelper.EXCEPT + "UserJoin1", gormhelper.EXCEPT + "UserPreload1"}, ","),
			joinStrings:      []string{gormhelper.EXCEPT + "UserJoin1", gormhelper.EXCEPT + "UserPreload1"},
			expectedJoins:    []string{"UserJoin2"},
			expectedPreloads: []string{"UserPreload2"},
		},
		{
			name:             strings.Join([]string{"UserJoin2", "UserPreload2"}, ","),
			joinStrings:      []string{"UserJoin2", "UserPreload2"},
			expectedJoins:    []string{"UserJoin2"},
			expectedPreloads: []string{"UserPreload2"},
		},
		{
			name:             strings.Join([]string{}, ","),
			joinStrings:      []string{},
			expectedJoins:    []string{"UserJoin1", "UserJoin2"},
			expectedPreloads: []string{"UserPreload1", "UserPreload2"},
		},
	}

	for _, tt := range tests {
		db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

		t.Run(tt.name, func(t *testing.T) {
			joinMapCopy := map[string][]interface{}{}
			preloadMapCopy := map[string][]interface{}{}
			maphelper.CopyMap(joinMap, &joinMapCopy)
			maphelper.CopyMap(preloadMap, &preloadMapCopy)

			db = gormhelper.Join(db, joinMapCopy, preloadMapCopy, tt.joinStrings...)

			joinStrs := []string{}
			for _, joinStatements := range db.Statement.Joins {
				joinStrs = append(joinStrs, joinStatements.Name)
			}
			if !arrayhelper.ValueEqual(joinStrs, tt.expectedJoins) {
				t.Errorf("joinStrings: %#v, Expected join %#v, found in SQL: %#v", tt.joinStrings, tt.expectedJoins, joinStrs)
			}

			preloadStrs := []string{}
			for preloadName := range db.Statement.Preloads {
				preloadStrs = append(preloadStrs, preloadName)
			}
			if !arrayhelper.ValueEqual(preloadStrs, tt.expectedPreloads) {
				t.Errorf("joinStrings: %#v, Expected preload %#v, found in SQL: %#v", tt.joinStrings, tt.expectedPreloads, preloadStrs)
			}
		})
	}
}
