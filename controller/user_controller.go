package controller

import (
	"github.com/danilocordeirodev/go-basics/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	seviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: seviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)

	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}