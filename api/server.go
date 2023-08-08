package api

import (
	"fmt"
	"log"

	"github.com/ayowilfred95/Golang-RESTful-API/controller"
	"github.com/ayowilfred95/Golang-RESTful-API/database"
	"github.com/ayowilfred95/Golang-RESTful-API/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ConnectDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Logs{})
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func Server(){
	router :=gin.Default()

	Routes := router.Group("/auth")
	Routes.POST("/register", controller.Register)
	Routes.POST("/login", controller.Login)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
