package model

import (

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init findUser model",
		zap.String("journey", "findUser"),
	)

	return nil, nil
} 