package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wendryuslima/user-manager/src/configuration/database/mongodb"
	"github.com/wendryuslima/user-manager/src/controller"
	"github.com/wendryuslima/user-manager/src/controller/routes"
	"github.com/wendryuslima/user-manager/src/model/repository"
	"github.com/wendryuslima/user-manager/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect database, error",
			err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	fmt.Println(os.Getenv("DB_HOST"))

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
