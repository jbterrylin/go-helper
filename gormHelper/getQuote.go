package gormhelper

import "gorm.io/gorm"

func GetQuote(db *gorm.DB) string {
	dialect := db.Dialector.Name()
	switch dialect {
	case MYSQL:
		return "`"
	case POSTGRES:
		return "\""
	case SQLITE:
		return "`"
	default:
		return "`"
	}
}
