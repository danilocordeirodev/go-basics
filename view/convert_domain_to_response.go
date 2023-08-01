package view

import (
	"github.com/danilocordeirodev/go-basics/controller/dto/response"
	"github.com/danilocordeirodev/go-basics/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
		return response.UserResponse{
			ID:    userDomain.GetID(),
			Email: userDomain.GetEmail(),
			Name:  userDomain.GetName(),
			Age:   userDomain.GetAge(),
		}
	
}