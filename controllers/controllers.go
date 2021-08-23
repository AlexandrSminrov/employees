package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/AlexandrSminrov/employees/models"
	"github.com/AlexandrSminrov/employees/repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sort"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req GET ALL ip: %v\n", r.RemoteAddr)

	sortBy := r.URL.Query().Get("sortBy")

	st, err := repositories.ConnDb.GetAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := fmt.Fprint(w, "Request error, repeat the ride"); err != nil {
			log.Println(err)

		}
		return
	}

	if sortBy == "idDown" {
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
	if _, err := w.Write(res); err != nil {
		log.Println(err)
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
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

	id, err := repositories.ConnDb.AddEmployee(st, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, id)

}

func GetByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req GetByID ip: %v\n", r.RemoteAddr)
	vars := mux.Vars(r)

	res, err := repositories.ConnDb.GetByID(vars["id"], r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req UpEmployee ip: %v\n", r.RemoteAddr)
	vars := mux.Vars(r)

	st := &models.DbStruct{}

	if err := json.NewDecoder(r.Body).Decode(st); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "JSON ERROR")
		log.Println("JSON ERROR", err)
		return
	}

	if err := repositories.ConnDb.UpEmploee(vars["id"], st, r.Context()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, "ok")

}
