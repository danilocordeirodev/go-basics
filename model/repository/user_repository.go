package repository

import (
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository (
	database *mongo.Database,
) UserRepository {
	return &userRepository {
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) ( model.UserDomainInterface, *rest_err.RestErr)
}