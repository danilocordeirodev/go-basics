package model

import (
	"fmt"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model",
		zap.String("journey", "createUser"),
	)

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}