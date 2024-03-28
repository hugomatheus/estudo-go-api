package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbTest *gorm.DB

func SetupDBTest(entities ...interface{}) {
	var err error
	dbTest, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbTest.AutoMigrate(entities...)
}

func GetDBTest() *gorm.DB {
	return dbTest
}

func CloseDBTest() {
	sqlDB, err := dbTest.DB()
	if err != nil {
		panic("failed to get sqldb")
	}

	err = sqlDB.Close()
	if err != nil {
		panic("failed to close database")
	}
}
