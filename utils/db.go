package utils

import (
	"fmt"

	"github.com/pxwxnvermx/todo-rest/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db.Debug()

	if err != nil {
		fmt.Print(err)
	}

	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		fmt.Print(err)
	}

	return db, nil
}
