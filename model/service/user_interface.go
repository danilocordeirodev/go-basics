package service

import (
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"github.com/danilocordeirodev/go-basics/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}


type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserById(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}