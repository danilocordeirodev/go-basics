package service

import (
	"fmt"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model",
		zap.String("journey", "createUser"),
	)

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}