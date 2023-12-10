package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedro-phd/first-api-go/src/configuration/logger"
	"github.com/pedro-phd/first-api-go/src/controller/routes"
)

func main() {

	// Loading .env file
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Info(fmt.Sprintf("🚀 System online in %s ambient", os.Getenv("AMBIENT")))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
