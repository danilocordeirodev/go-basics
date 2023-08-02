package controller

import (
	"net/http"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/config/validation"
	"github.com/danilocordeirodev/go-basics/controller/dto/request"
	"github.com/danilocordeirodev/go-basics/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func  (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
	
		logger.Error("Error trying to validate user info", err,
				zap.String("journey", "updateUser"),
			)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to UpdateUser controller - UserID is not a valid identifier", err,
				zap.String("journey", "UpdateUser"),
			)
		errRest := rest_err.NewBadRequestError("UserID is not a valid identifier")
		c.JSON(errRest.Code, errRest)
	}


	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)


	err := uc.service.UpdateUser(userId, domain) 
	
	if err != nil {
		logger.Error("Error trying to call updateUser controller",
		err,
		zap.String("journey", "updateUser"),
	)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)

}