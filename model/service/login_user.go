package service

import (
	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init loginUser model",
		zap.String("journey", "loginUser"),
	)

	userDomain.EncryptPassword()

	user, err :=ud.findUserByEmailAndPassword(userDomain.GetEmail(), userDomain.GetPassword())
	
	if err != nil {
		return nil, err
	}

	logger.Info("Init loginUser executed successfully",
		zap.String("journey", "loginUser"),
	)

	return user, nil
}