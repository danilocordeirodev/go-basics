package repository

import (
	"context"

	"os"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.uber.org/zap"
)



func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init DeleteUser repository", zap.String("journey", "DeleteUser"))
	
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}


	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		logger.Error("Error trying to DeleteUser repository", err,
				zap.String("journey", "DeleteUser"),
			)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("Init DeleteUser repository executed successfully", 
		zap.String("userId", userId),
		zap.String("journey", "DeleteUser"),
	)
	
	return nil
	
}