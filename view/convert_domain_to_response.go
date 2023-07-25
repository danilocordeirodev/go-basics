package view

import (
	"github.com/danilocordeirodev/go-basics/controller/dto/response"
	"github.com/danilocordeirodev/go-basics/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
		return response.UserResponse{
			ID:    "test",
			Email: userDomain.GetEmail(),
			Name:  userDomain.GetName(),
			Age:   userDomain.GetAge(),
		}
	
}