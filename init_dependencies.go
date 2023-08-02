package main

import (
	"github.com/danilocordeirodev/go-basics/controller"
	"github.com/danilocordeirodev/go-basics/model/repository"
	"github.com/danilocordeirodev/go-basics/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}