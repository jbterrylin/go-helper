package gormhelper_test

import (
	"reflect"
	"testing"

	gormhelper "github.com/jbterrylin/go-helper/gormHelper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestSelect(t *testing.T) {
	type User struct {
		ID       uint
		Name     string
		Age      int
		Salary   float64
		Location string
	}

	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	db.AutoMigrate(User{})

	tests := []struct {
		name             string
		model            interface{}
		oriTableName     string
		oriPreFix        string
		sumList          []string
		ignoreList       []string
		targetTableName  string
		targetPrefix     string
		targetPostfix    string
		customList       []string
		needSort         bool
		expectedSQLParts []string
	}{
		{
			name:            "basic select",
			model:           &User{},
			oriTableName:    "",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "",
			targetPrefix:    "",
			targetPostfix:   "",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`id`",
				"`name`",
				"`age`",
				"`salary`",
				"`location`",
			},
		},
		{
			name:            "select with same table name",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "users",
			targetPrefix:    "",
			targetPostfix:   "",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id`",
				"`users`.`name`",
				"`users`.`age`",
				"`users`.`salary`",
				"`users`.`location`",
			},
		},
		{
			name:            "rename table",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "users1",
			targetPrefix:    "",
			targetPostfix:   "",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id` AS `users1`.`id`",
				"`users`.`name` AS `users1`.`name`",
				"`users`.`age` AS `users1`.`age`",
				"`users`.`salary` AS `users1`.`salary`",
				"`users`.`location` AS `users1`.`location`",
			},
		},
		{
			name:            "ignore field",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{"age"},
			targetTableName: "users",
			targetPrefix:    "",
			targetPostfix:   "",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id`",
				"`users`.`name`",
				"`users`.`salary`",
				"`users`.`location`",
			},
		},
		{
			name:            "sum field",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{"salary"},
			ignoreList:      []string{},
			targetTableName: "users",
			targetPrefix:    "",
			targetPostfix:   "",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id`",
				"`users`.`name`",
				"`users`.`age`",
				"SUM(`users`.`salary`) AS `users`.`salary`",
				"`users`.`location`",
			},
		},
		{
			name:            "custom select",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "users",
			targetPrefix:    "",
			targetPostfix:   "",
			customList:      []string{"SUM(`users`.`salary`) AS `users`.`salary`"},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id`",
				"`users`.`name`",
				"`users`.`age`",
				"`users`.`salary`",
				"`users`.`location`",
				"SUM(`users`.`salary`) AS `users`.`salary`",
			},
		},
		{
			name:            "target prefix",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "users",
			targetPrefix:    "pre__",
			targetPostfix:   "",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id` AS `users`.`pre__id`",
				"`users`.`name` AS `users`.`pre__name`",
				"`users`.`age` AS `users`.`pre__age`",
				"`users`.`salary` AS `users`.`pre__salary`",
				"`users`.`location` AS `users`.`pre__location`",
			},
		},
		{
			name:            "target post",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "users",
			targetPrefix:    "",
			targetPostfix:   "__post",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id` AS `users`.`id__post`",
				"`users`.`name` AS `users`.`name__post`",
				"`users`.`age` AS `users`.`age__post`",
				"`users`.`salary` AS `users`.`salary__post`",
				"`users`.`location` AS `users`.`location__post`",
			},
		},
		{
			name:            "target postfix2",
			model:           &User{},
			oriTableName:    "users",
			oriPreFix:       "",
			sumList:         []string{},
			ignoreList:      []string{},
			targetTableName: "",
			targetPrefix:    "",
			targetPostfix:   "__post",
			customList:      []string{},
			needSort:        false,
			expectedSQLParts: []string{
				"`users`.`id` AS `id__post`",
				"`users`.`name` AS `name__post`",
				"`users`.`age` AS `age__post`",
				"`users`.`salary` AS `salary__post`",
				"`users`.`location` AS `location__post`",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlParts := gormhelper.Select(db, tt.model, tt.oriTableName, tt.oriPreFix, tt.sumList, tt.ignoreList, tt.targetTableName, tt.targetPrefix, tt.targetPostfix, tt.customList, tt.needSort)
			if !reflect.DeepEqual(sqlParts, tt.expectedSQLParts) {
				t.Errorf("Select() = %#v, want %#v", sqlParts, tt.expectedSQLParts)
			}
		})
	}
}
