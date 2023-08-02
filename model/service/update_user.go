package service

import (
	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface)  *rest_err.RestErr {
	logger.Info("Init updateUser service",
		zap.String("journey", "updateUser"),
	)


	err :=ud.userRepository.UpdateUser(userId, userDomain)
	
	if err != nil {
		logger.Error("Error trying to call repository", err,
		zap.String("journey", "updateUser"),
	)
		return err
	}

	logger.Info("updateUser service executed successfully",
		zap.String("journey", "updateUser"),
	)

	return nil

}