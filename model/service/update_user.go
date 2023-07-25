package service

import (
	"fmt"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model",
		zap.String("journey", "updateUser"),
	)

	userDomain.EncryptPassword()

	fmt.Println(userDomain)

	return nil
}