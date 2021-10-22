package router

import (
	"go-api-tut/pkg/api/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/grpc/hello", user.Hello)

	r.GET("/users", user.GetUsers)
	r.POST("/users", user.CreateUser)
	r.GET("/users/:id", user.GetUser)
	r.PUT("/users/:id", user.UpdateUser)
	r.DELETE("/users/:id", user.DeleteUser)

	return r
}
