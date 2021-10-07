package core

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		users := []User{}
		Db.Find(&users)

		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		user := User{Name: c.PostForm("name")}
		Db.Create(&user)

		c.JSON(http.StatusCreated, user)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		user := User{}
		Db.First(&user, c.Param("id"))

		c.JSON(http.StatusOK, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		user := User{}
		Db.First(&user, c.Param("id"))

		err := c.BindJSON(&user)
		if err != nil {
			return
		}

		Db.Save(&user)

		c.JSON(http.StatusOK, user)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		Db.Delete(&User{}, c.Param("id"))

		c.JSON(http.StatusNoContent, nil)
	})

	return r
}
