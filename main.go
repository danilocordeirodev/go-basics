package main

import (
	"context"
	"log"

	"github.com/danilocordeirodev/go-basics/config/database/mongodb"
	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/controller"
	"github.com/danilocordeirodev/go-basics/controller/routes"
	"github.com/danilocordeirodev/go-basics/model/repository"
	"github.com/danilocordeirodev/go-basics/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	logger.Info("Starting server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error creating database connection, err=%s \n", err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
