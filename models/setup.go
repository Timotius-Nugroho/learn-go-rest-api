package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/go_rest_api_gin"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{}, &Category{})
	DB = database

}
