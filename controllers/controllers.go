package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/AlexandrSminrov/employees/models"
	"github.com/gorilla/mux"
	//_ "github.com/lib/pq"
)

type server struct {
	dbClient models.DBClient
}

// NewServer initializes the service configuration
func NewServer(dbClient *models.DBClient) models.Server {
	return &server{
		dbClient: *dbClient,
	}
}

// GetAll will return all records from the table
// @Summary return all records from the table
// @Tags employee
// @Description return all records from the table "employee"
// @ID GetAll
// @Router /employee [get]
// @Success 200 {array} models.DbStruct{}
// @Failure 500 "string: error"
func (server *server) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req GET ALL ip: %v\n", r.RemoteAddr)

	sortBy := r.URL.Query().Get("sortBy")

	st, err := server.dbClient.GetAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = fmt.Fprint(w, "Request error, repeat the ride"); err != nil {
			log.Println(err)
		}
		return
	}

	if sortBy == "idUP" {
		sort.Slice(st, func(i, j int) (less bool) {
			return st[i].ID > st[j].ID
		})
	}

	res, err := json.Marshal(st)
	if err != nil {
		log.Println("Marshal json error")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(res); err != nil {
		log.Println(err)
	}
}

// AddEmployee add record to table
// @Summary add employee
// @Tags employee
// @Description add employee from the table "employee"
// @Accept json
// @Produce json
// @ID AddEmployee
// @Router /employee [post]
// @Param input body models.DbStruct{} true "Body request"
// @Success 201 "intenger id"
// @Failure 400 "string error json or is not valid date"
// Failure 500 {error}
func (server *server) AddEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req AddEmployee ip: %v\n", r.RemoteAddr)

	st := &models.DbStruct{}

	if err := json.NewDecoder(r.Body).Decode(st); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "JSON ERROR")
		log.Println("JSON ERROR", err)
		return
	}

	if err := st.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	id, err := server.dbClient.AddEmployee(st, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, id)
}

// GetByID get employee by id
// @Summary get employee
// @Tags employee
// @Description get employee by id from the table "employee"
// @ID GetByID
// @Router /employee/{employeeID} [get]
// @Success 200 {object} models.DbStruct{}
// Failure 500 "string error"
func (server *server) GetByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req GetByID ip: %v\n", r.RemoteAddr)
	vars := mux.Vars(r)

	res, err := server.dbClient.GetByID(vars["id"], r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpEmployee update employee by id
// @Summary update employee
// @Tags employee
// @Description update employee by id from the table "employee"
// @ID UpEmployee
// @Router /employee/{employeeID} [put]
// @Success 200 {object} models.DbStruct{}
// Failure 400 "string error"
// Failure 500 "string error"
func (server *server) UpEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req UpEmployee ip: %v\n", r.RemoteAddr)
	vars := mux.Vars(r)

	st := &models.DbStruct{}

	if err := json.NewDecoder(r.Body).Decode(st); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "JSON ERROR")
		log.Println("JSON ERROR", err)
		return
	}

	if err := server.dbClient.UpEmployee(vars["id"], st, r.Context()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "ok")
}
