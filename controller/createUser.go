package controller

import (
	"fmt"
	"net/http"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/validation"
	"github.com/danilocordeirodev/go-basics/controller/dto/request"
	"github.com/danilocordeirodev/go-basics/controller/dto/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		

	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
	
		logger.Error("Error trying to create user", err,
			)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

	response := response.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("User created successfully",
		

	)

	c.JSON(http.StatusOK, response)

}
