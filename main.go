package main

import (
	"github.com/ayowilfred95/Golang-RESTful-API/api"
)

func main() {
	api.LoadEnv()
	api.ConnectDatabase()
	api.Server()
}



