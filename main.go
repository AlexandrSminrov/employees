package main

import (
	"log"

	"github.com/AlexandrSminrov/employees/bootstrap"
	"github.com/AlexandrSminrov/employees/configs"
)

// @title Employees
// @version 1.0
// @description employee base management
// @host localhost:8080
// @BasePath /
func main() {
	config, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := bootstrap.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}

	bootstrap.InitServer(dbClient, config)
}
