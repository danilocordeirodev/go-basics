package service

import (
	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("Init deleteUser service",
		zap.String("journey", "deleteUser"),
	)


	err := ud.userRepository.DeleteUser(userId)
	
	if err != nil {
		logger.Error("Error trying to call repository", err,
		zap.String("journey", "deleteUser"),
	)
		return err
	}

	logger.Info("deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}