package main

import (
	"github.com/AlexandrSminrov/employees/repositories"
	"github.com/AlexandrSminrov/employees/routers"
	"log"
	"net/http"
)

func main() {
	//if err := bootstrap.InitConfig(); err != nil {
	//	log.Fatal(err)
	//}

	if err := repositories.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Handler: routers.GetRoutes(),
		Addr:    ":8080",
	}

	log.Println("start")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
