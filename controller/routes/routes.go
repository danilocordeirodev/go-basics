package routes

import (
	"github.com/danilocordeirodev/go-basics/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface)  {
	r.GET("/users/:userId", userController.FindUserByID)
	r.GET("/users/by-email/:userEmail", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)
}