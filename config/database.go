package config

import (
	"github.com/athun/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Dbconnect() (DB *gorm.DB) {
	dsn := "root:@tcp(127.0.0.1:3306)/webapp?parseTime=true"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.User{})
	return DB
}
