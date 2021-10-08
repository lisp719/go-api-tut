package core

import (
	"errors"
	"go-api-tut/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		users := []models.User{}
		Db.Find(&users)

		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		user := models.User{}

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		Db.Create(&user)

		c.JSON(http.StatusCreated, user)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		user := models.User{}
		result := Db.First(&user, c.Param("id"))

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		user := models.User{}
		result := Db.First(&user, c.Param("id"))

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		Db.Save(&user)

		c.JSON(http.StatusOK, user)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		user := models.User{}
		result := Db.First(&user, c.Param("id"))

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		Db.Delete(&user)

		c.JSON(http.StatusNoContent, nil)
	})

	return r
}
