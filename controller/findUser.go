package controller

import (
	"net/http"
	"net/mail"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func  (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserById controller", zap.String("journey", "FindUserById"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying validate userID", err, zap.String("journey", "FindUserById"))
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid user ID",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserById(userId)
	if err != nil {
		logger.Error("Error trying to call FindUserById services", err, zap.String("journey", "FindUserById"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed successfully", 
		zap.String("journey", "FindUserById"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func  (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying validate userEmail", err, zap.String("journey", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail services", err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", 
		zap.String("journey", "FindUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}