package controller

import (
	"net/http"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func  (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)


	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to deleteUser controller - UserID is not a valid identifier", err,
				zap.String("journey", "deleteUser"),
			)
		errRest := rest_err.NewBadRequestError("UserID is not a valid identifier")
		c.JSON(errRest.Code, errRest)
	}


	err := uc.service.DeleteUser(userId) 
	
	if err != nil {
		logger.Error("Error trying to call deleteUser controller",
		err,
		zap.String("journey", "deleteUser"),
	)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusNoContent)

}