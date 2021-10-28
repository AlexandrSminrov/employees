package routers

import (
	"net/http"

	_ "github.com/AlexandrSminrov/employees/docs"
	"github.com/AlexandrSminrov/employees/models"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

type route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func GetRoutes(s models.Server) *mux.Router {
	routes := []route{
		{
			"AllEmployees",
			http.MethodGet,
			"/employee",
			s.GetAll,
		},
		{
			"AddEmployee",
			http.MethodPost,
			"/employee",
			s.AddEmployee,
		},
		{
			"GetEmployee",
			http.MethodGet,
			"/employee/{id:[0-9]+}",
			s.GetByID,
		},
		{
			"UpEmployee",
			http.MethodPut,
			"/employee/{id:[0-9]+}",
			s.UpEmployee,
		},
	}

	m := mux.NewRouter()

	for _, r := range routes {
		m.Methods(r.Method).Name(r.Name).Path(r.Path).Handler(r.HandlerFunc)
	}

	// swagger
	m.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return m
}
