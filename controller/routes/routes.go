package routes

import (
	"github.com/danilocordeirodev/go-basics/controller"
	"github.com/danilocordeirodev/go-basics/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface)  {
	r.GET("/users/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/users/by-email/:userEmail",  model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId",  model.VerifyTokenMiddleware,userController.UpdateUser)
	r.DELETE("/users/:userId",  model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)
}