package controllers

import (
	"errors"
	"go-api-tut/core"
	"go-api-tut/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}

	if q := c.Query("q"); q == "" {
		core.Db.Find(&users)
	} else {
		core.Db.Where("name LIKE ?", "%"+q+"%").Find(&users)
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	core.Db.Create(&user)

	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	user := models.User{}
	result := core.Db.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	user := models.User{}
	result := core.Db.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	core.Db.Save(&user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	user := models.User{}
	result := core.Db.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	core.Db.Delete(&user)

	c.JSON(http.StatusNoContent, nil)
}
