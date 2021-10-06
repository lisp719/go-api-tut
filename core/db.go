package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

var Db *gorm.DB

func SetupDb() {
	var err error
	dsn := "root:@tcp(mysql:3306)/go_api_tut?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
}
