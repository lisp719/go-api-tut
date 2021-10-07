package core

import (
	"os"

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
	dsn := os.Getenv("GORM_DSN")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
}
