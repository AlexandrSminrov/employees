package models

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

// DBClient methods db
type DBClient interface {
	ConnectDB() error
	GetAll(ctx context.Context) ([]*DbStruct, error)
	AddEmployee(ctx context.Context, dbStruct *DbStruct) (int, error)
	GetByID(ctx context.Context, id string) ([]byte, error)
	UpEmployee(ctx context.Context, id string, st *DbStruct) error
}

// Server handle server
type Server interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	AddEmployee(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	UpEmployee(w http.ResponseWriter, r *http.Request)
}

// regexp patterns
var (
	onlyRu    = regexp.MustCompile(`[^А-Яа-я]`)
	onlyRuEng = regexp.MustCompile(`[^A-Za-zА-Яа-я]`)
	onlyNum   = regexp.MustCompile(`[^0-9]`)
	address   = regexp.MustCompile(`[^а-яА-Я0-9,.\s№]`)
	email     = regexp.MustCompile(`[^а-яА-Я0-9,.\s№]`)
)

// DbStruct base structure
type DbStruct struct {
	ID          int    `json:"id,omitempty"`
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	MiddleName  string `json:"middlename,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Address     string `json:"addres,omitempty"`
	Department  string `json:"department,omitempty"`
	AboutMe     string `json:"aboutme,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
}

// Validate verifies the request
func (st *DbStruct) Validate() error {
	if onlyRu.MatchString(st.FirstName) {
		return fmt.Errorf("FirstName ERROR")
	}

	if onlyRu.MatchString(st.LastName) {
		return fmt.Errorf("LastName ERROR")
	}

	if onlyRu.MatchString(st.MiddleName) {
		return fmt.Errorf("MiddleName ERROR")
	}

	if _, err := time.Parse("01.02.2006", st.DateOfBirth); err != nil && len(st.DateOfBirth) > 1 {
		return fmt.Errorf("date ERROR ")
	}

	if address.MatchString(st.Address) {
		return fmt.Errorf("address ERROR ")
	}

	if onlyRuEng.MatchString(st.Department) {
		return fmt.Errorf("department ERROR ")
	}

	if address.MatchString(st.AboutMe) {
		return fmt.Errorf("AboutMe ERROR")
	}

	if onlyNum.MatchString(st.Phone) {
		return fmt.Errorf("phone number ERROR ")
	}

	if !email.MatchString(st.Email) && len(st.Email) > 1 {
		return fmt.Errorf("email ERROR ")
	}

	return nil
}
