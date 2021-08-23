package routers

import (
	"github.com/AlexandrSminrov/employees/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func GetRoutes() *mux.Router {
	routes := []route{
		{
			"AllEmployees",
			"GET",
			"/employee",
			controllers.GetAll,
		},
		{
			"AddEmployee",
			"POST",
			"/employee",
			controllers.AddEmployee,
		},
		{
			"GetEmployee",
			"GET",
			"/employee/{id:[0-9]+}",
			controllers.GetByID,
		},
		{
			"UpEmployee",
			"PUT",
			"/employee/{id:[0-9]+}",
			controllers.UpEmployee,
		},
	}

	m := mux.NewRouter()

	for _, r := range routes {
		m.Methods(r.Method).Name(r.Name).Path(r.Path).Handler(r.HandlerFunc)
	}

	return m

}
