package model

import (
	"fmt"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser(string) *rest_err.RestErr {
	logger.Info("Init updateUser model",
		zap.String("journey", "updateUser"),
	)

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}