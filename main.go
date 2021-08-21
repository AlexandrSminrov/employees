package main

import (
	"github.com/AlexandrSminrov/employees/bootstrap"

	"log"
)

func main() {
	if err := bootstrap.InitConfig(); err != nil {
		log.Fatal(err)
	}

	if err := bootstrap.InitConnectDB(); err != nil {
		log.Fatal(err)
	}

}
