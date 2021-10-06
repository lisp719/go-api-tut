package main

import (
	"go-api-tut/core"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		users := db.Find(&[]core.User{})

		c.JSON(200, gin.H{
			"count": users.RowsAffected,
		})
	})

	return r
}

func main() {
	db = core.DbConnect()

	r := setupRouter()
	r.Run(":8080")
}
