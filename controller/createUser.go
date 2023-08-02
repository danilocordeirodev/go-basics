package controller

import (
	"net/http"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/validation"
	"github.com/danilocordeirodev/go-basics/controller/dto/request"
	"github.com/danilocordeirodev/go-basics/model"
	"github.com/danilocordeirodev/go-basics/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserCreateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
	
		logger.Error("Error trying to validate user info", err,
				zap.String("journey", "createUser"),
			)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)


	domainResult, err := uc.service.CreateUser(domain) 
	
	if err != nil {
		logger.Error("Error trying to call CreateUser service",
		err,
		zap.String("journey", "createUser"),
	)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreatedUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
