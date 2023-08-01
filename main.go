package main

import (
	"log"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"github.com/danilocordeirodev/go-basics/controller"
	"github.com/danilocordeirodev/go-basics/controller/routes"
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
	
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
