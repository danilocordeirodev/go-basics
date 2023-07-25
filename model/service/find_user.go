package service

import (
	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/config/rest_err"
	"github.com/danilocordeirodev/go-basics/model"
	"go.uber.org/zap"
)

func (*userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser model",
		zap.String("journey", "findUser"),
	)

	return nil, nil
} 