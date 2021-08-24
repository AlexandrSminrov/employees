package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AlexandrSminrov/employees/models"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type DbQuery struct {
	DB *sql.DB
}

var ConnDb *DbQuery

func ConnectDB() error {
	dbStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"db", 5432,
		os.Getenv("pgUser"),
		os.Getenv("pgPass"),
		os.Getenv("pgDb"),
	)

	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		return fmt.Errorf("Connection error: %v ", err)
	}

	maxConns, err := strconv.Atoi(os.Getenv("MaxConns"))
	if err != nil {
		return fmt.Errorf("MaxConns read error: %v ", err)
	}

	idleConns, err := strconv.Atoi(os.Getenv("IdleConns"))
	if err != nil {
		return fmt.Errorf("IdleConns read error: %v ", err)
	}

	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(idleConns)

	ConnDb = &DbQuery{DB: db}

	return nil
}

func (db *DbQuery) GetAll(ctx context.Context) ([]models.DbStruct, error) {

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	var query = "SELECT * FROM public.emploees"

	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Query %v \nError: %v\n", query, err)
	}
	defer rows.Close()

	var structs []models.DbStruct

	for rows.Next() {
		var st models.DbStruct

		if err := rows.Scan(
			&st.ID,
			&st.FirstName,
			&st.LastName,
			&st.MiddleName,
			&st.BDate,
			&st.Address,
			&st.Department,
			&st.AboutMe,
			&st.Tnumber,
			&st.Email,
		); err != nil {
			return nil, err
		}

		st.Address = ""
		st.AboutMe = ""

		structs = append(structs, st)
	}

	return structs, nil
}

func (db *DbQuery) AddEmployee(dbStruct *models.DbStruct, ctx context.Context) (int, error) {

	query := "INSERT INTO public.emploees (firstname, lastname, middlename, bdate,addres, department, aboutMe, tnumber, email)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 ) RETURNING id"

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	rows, err := db.DB.QueryContext(ctx, query, dbStruct.FirstName,
		dbStruct.LastName,
		dbStruct.MiddleName,
		dbStruct.BDate,
		dbStruct.Address,
		dbStruct.Department,
		dbStruct.AboutMe,
		dbStruct.Tnumber,
		dbStruct.Email,
	)

	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("Row err: %v ", err)
		}
	}

	return id, err
}

func (db *DbQuery) GetByID(id string, ctx context.Context) ([]byte, error) {

	rows, err := db.DB.QueryContext(ctx, "select * from public.emploees where id in ($1)", id)
	if err != nil {
		return nil, err
	}

	var st models.DbStruct

	for rows.Next() {
		if err := rows.Scan(
			&st.ID,
			&st.FirstName,
			&st.LastName,
			&st.MiddleName,
			&st.BDate,
			&st.Address,
			&st.Department,
			&st.AboutMe,
			&st.Tnumber,
			&st.Email,
		); err != nil {
			return nil, err
		}
	}

	if st.FirstName == "" {
		return nil, fmt.Errorf("id:%v not fond", id)
	}

	res, err := json.Marshal(st)
	if err != nil {
		return nil, fmt.Errorf("Marshal error: %v\n", err)
	}

	return res, nil
}

func (db *DbQuery) UpEmploee(id string, st *models.DbStruct, ctx context.Context) error {

	query := "UPDATE public.emploees SET "

	v := reflect.ValueOf(*st)
	typeOfS := v.Type()

	for i := 1; i < v.NumField(); i++ {
		val := fmt.Sprint(v.Field(i).Interface())
		if len(val) > 0 {
			query += strings.ToLower(typeOfS.Field(i).Name) + "='" + fmt.Sprint(val) + "' "
		}
	}

	query += "WHERE id=$1"

	fmt.Println(query)
	_, err := db.DB.QueryContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
