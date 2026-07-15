package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wendryuslima/user-manager/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
	fmt.Println(os.Getenv("DB_HOST"))

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	router.Run(":8080")
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}

}