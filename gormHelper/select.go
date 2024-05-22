package gormhelper

import (
	"fmt"
	"sort"
	"strings"

	"gorm.io/gorm"
)

// SelectHelper for solve hard code select column
// e.x: want to remove 1 column only but need to write all column except no needed column
// e.x: want to sum 1 of the column but need to rewrite all of the column

func quoteText(quote, text string, addDot bool) string {
	if text == "" {
		return ""
	}
	dot := ""
	if addDot {
		dot = "."
	}
	return fmt.Sprintf("%s%s%s%s", quote, text, quote, dot)
}

func getBaseFieldName(text string) string {
	for _, as := range []string{" AS ", " as ", " As ", "aS"} {
		for strings.Contains(text, as) {
			text = strings.Split(text, as)[1]
		}
	}
	return text
}

// oriPreFix: can use for prefix ori table name (e.x: oriPreFix = "data_baidu" become select `data_baidu`.xx)
func Select(
	db *gorm.DB,
	model interface{},
	oriTableName string,
	oriPreFix string,
	sumList []string,
	ignoreList []string,
	targetTableName string,
	targetPrefix string,
	targetPostfix string,
	customList []string,
	needSort bool) []string {
	quote := GetQuote(db)
	columnTypes, err := db.Migrator().ColumnTypes(model)
	if err != nil {
		panic(err)
	}

	ignoreMap := make(map[string]bool, len(ignoreList))
	for _, columnName := range ignoreList {
		ignoreMap[columnName] = true
	}

	sumMap := make(map[string]bool, len(sumList))
	for _, sum := range sumList {
		sumMap[sum] = true
	}

	result := make([]string, 0)
	for _, column := range columnTypes {
		columnName := column.Name()

		if ignoreMap[columnName] {
			continue
		}

		isSumField := sumMap[columnName]

		oldColumnName := fmt.Sprintf("%s%s", quoteText(quote, oriTableName, true), quoteText(quote, oriPreFix+columnName, false))
		newColumnName := fmt.Sprintf("%s%s", quoteText(quote, targetTableName, true), quoteText(quote, targetPrefix+columnName+targetPostfix, false))

		selectField := ""
		if isSumField {
			selectField = fmt.Sprintf("SUM(%s) AS %s", oldColumnName, newColumnName)
		} else {
			if oldColumnName == newColumnName {
				selectField = oldColumnName
			} else {
				selectField = fmt.Sprintf("%s AS %s", oldColumnName, newColumnName)
			}
		}

		result = append(result, selectField)
	}

	if customList != nil {
		result = append(result, customList...)
	}

	if needSort {
		sort.Slice(result, func(i, j int) bool {
			return getBaseFieldName(result[i]) > getBaseFieldName(result[j])
		})
	}

	return result
}
