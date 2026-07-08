package usermanager

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
	fmt.Println(os.Getenv("DB_HOST"))

}