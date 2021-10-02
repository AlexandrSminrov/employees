package models

import (
	"fmt"
	"regexp"
	"time"
)

type Config struct {
	Application struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	Database struct {
		Type         string
		Host         string
		Port         int
		User         string `config:"envVar"`
		Password     string `config:"envVar"`
		Dbname       string `config:"envVar"`
		MaxIdleConns int
		MaxOpenConns int
	}
}

type DbStruct struct {
	ID         int    `json:"id,omitempty"`
	FirstName  string `json:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	MiddleName string `json:"middlename,omitempty"`
	BDate      string `json:"bdate,omitempty"`
	Address    string `json:"addres,omitempty"`
	Department string `json:"department,omitempty"`
	AboutMe    string `json:"aboutme,omitempty"`
	Tnumber    string `json:"tnumber,omitempty"`
	Email      string `json:"email,omitempty"`
}

func (st *DbStruct) Validate() error {

	flm := regexp.MustCompile(`[^А-Яа-я]`)
	dep := regexp.MustCompile(`[^A-Za-zА-Яа-я]`)
	phone := regexp.MustCompile(`[^0-9]`)
	addres := regexp.MustCompile(`[^а-яА-Я0-9,.\s№]`)
	email := regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)

	if flm.MatchString(st.FirstName) {
		return fmt.Errorf("FirstName ERROR")
	}

	if flm.MatchString(st.LastName) {
		return fmt.Errorf("LastName ERROR")
	}

	if flm.MatchString(st.MiddleName) {
		return fmt.Errorf("MiddleName ERROR")
	}

	if _, err := time.Parse("01.02.2006", st.BDate); err != nil && len(st.BDate) > 1 {
		return fmt.Errorf("Date ERROR ")
	}

	if addres.MatchString(st.Address) {
		return fmt.Errorf("Address ERROR ")
	}

	if dep.MatchString(st.Department) {
		return fmt.Errorf("Department ERROR ")
	}

	if addres.MatchString(st.AboutMe) {
		return fmt.Errorf("AboutMe ERROR")
	}

	if phone.MatchString(st.Tnumber) {
		return fmt.Errorf("Phone number ERROR ")
	}

	if !email.MatchString(st.Email) && len(st.Email) > 1 {
		return fmt.Errorf("Email ERROR ")
	}

	return nil

}
