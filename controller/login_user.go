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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)

	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
	
		logger.Error("Error trying to validate user info", err,
				zap.String("journey", "loginUser"),
			)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)


	domainResult, err := uc.service.LoginUser(domain) 
	
	if err != nil {
		logger.Error("Error trying to call loginUser service",
		err,
		zap.String("journey", "loginUser"),
	)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("logindUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}