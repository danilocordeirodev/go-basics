package routes

import (
	"github.com/danilocordeirodev/go-basics/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/:userId", controller.FindUserByID)
	r.GET("/users/by-email/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("/users/:userId", controller.DeleteUser)
}