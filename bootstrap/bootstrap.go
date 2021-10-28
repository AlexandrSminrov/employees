package bootstrap

import (
	"log"
	"net/http"

	"github.com/AlexandrSminrov/employees/configs"
	"github.com/AlexandrSminrov/employees/controllers"
	"github.com/AlexandrSminrov/employees/models"
	"github.com/AlexandrSminrov/employees/repositories"
	"github.com/AlexandrSminrov/employees/routers"
)

// InitDB init postgres
func InitDB(config *configs.ServerConfig) (*models.DBClient, error) {
	dbClient := repositories.NewDBClient(&config.DBConfig)
	return &dbClient, dbClient.ConnectDB()
}

// InitServer server
func InitServer(dbClient *models.DBClient, config *configs.ServerConfig) {
	server := controllers.NewServer(dbClient)
	srv := &http.Server{
		Handler: routers.GetRoutes(server),
		Addr:    ":" + config.Port,
	}

	log.Println("start service...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
