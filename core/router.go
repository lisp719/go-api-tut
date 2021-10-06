package core

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		users := Db.Find(&[]User{})

		c.JSON(200, gin.H{
			"count": users.RowsAffected,
		})
	})

	return r
}
