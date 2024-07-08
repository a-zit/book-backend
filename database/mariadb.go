package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMariaDB(uri string) *sql.DB {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDB
}
