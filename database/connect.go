package database

import (
	"fmt"

	"github.com/Satr10/fufufafa-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./fufufafa.db"), &gorm.Config{})
	if err != nil {
		panic("gagal connect ke database")
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Post{})
	fmt.Println("Database Migrated")
}
