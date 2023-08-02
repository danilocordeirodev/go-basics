package service

import (
	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.uber.org/zap"
)


func (ud *userDomainService) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail services", zap.String("journey", "FindUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserById(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById services", zap.String("journey", "FindUserById"))
	return ud.userRepository.FindUserById(id)
}