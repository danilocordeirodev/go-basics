package repository

import (
	"context"
	"os"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"github.com/danilocordeirodev/go-basics/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)



func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))
	
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		logger.Error("Error trying to createUser repository", err,
				zap.String("journey", "createUser"),
			)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID =  result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
	
}