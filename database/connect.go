package database

import (
	"fmt"

	"github.com/Satr10/fufufafa-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require TimeZone=Asia/Jakarta", config.Config("host"), config.Config("user"), config.Config("password"), config.Config("dbname"), config.Config("port"))
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		panic("gagal connect ke database")
	}
	fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&model.Post{})
	fmt.Println("Database Migrated")
}
