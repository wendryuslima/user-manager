package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wendryuslima/user-manager/src/controller"
	"github.com/wendryuslima/user-manager/src/controller/routes"
	"github.com/wendryuslima/user-manager/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println(os.Getenv("DB_HOST"))
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
