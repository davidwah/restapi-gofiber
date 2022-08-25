package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:password@tcp(172.18.0.2:3306)/restapi-fiber"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	DB = db
}
