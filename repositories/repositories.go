package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/AlexandrSminrov/employees/configs"
	"github.com/AlexandrSminrov/employees/models"
	_ "github.com/jackc/pgx/v4/stdlib" // driver pgx
)

// dbClient client model
type dbClient struct {
	DB     *sql.DB
	config *configs.DBConfig
}

// NewDBClient ...
func NewDBClient(config *configs.DBConfig) models.DBClient {
	return &dbClient{
		config: config,
	}
}

// ConnectDB the function of connecting to the database
func (db *dbClient) ConnectDB() error {
	if db.DB != nil {
		db.DB.Close()
	}

	dbStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.config.Host,
		db.config.Port,
		db.config.User,
		db.config.Password,
		db.config.DBName,
	)

	var err error

	db.DB, err = sql.Open("pgx", dbStr)
	if err != nil {
		return fmt.Errorf("connection error: %v ", err)
	}

	db.DB.SetMaxOpenConns(db.config.MaxConn)
	db.DB.SetMaxIdleConns(db.config.MaxIdleConn)
	db.DB.SetConnMaxIdleTime(db.config.TimeIdleConn)

	return nil
}

//
func (db *dbClient) GetAll(ctx context.Context) ([]*models.DbStruct, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	query := "SELECT id, firstname, lastname, middlename, date_of_birth,addres, department, about_me, phone, email  FROM public.emploees"

	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query %v\nerror: %v", query, err)
	}
	defer rows.Close()

	var structs []*models.DbStruct

	for rows.Next() {
		var st models.DbStruct

		if err = rows.Scan(
			&st.ID,
			&st.FirstName,
			&st.LastName,
			&st.MiddleName,
			&st.DateOfBirth,
			&st.Address,
			&st.Department,
			&st.AboutMe,
			&st.Phone,
			&st.Email,
		); err != nil {
			return nil, err
		}

		st.Address = ""
		st.AboutMe = ""

		structs = append(structs, &st)
	}

	return structs, nil
}

// AddEmployee ...
func (db *dbClient) AddEmployee(ctx context.Context, dbStruct *models.DbStruct) (int, error) {
	query := "INSERT INTO public.emploees (firstname, lastname, middlename, bdate,addres, department, aboutMe, tnumber, email)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 ) RETURNING id"

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	rows, err := db.DB.QueryContext(ctx, query, dbStruct.FirstName,
		dbStruct.LastName,
		dbStruct.MiddleName,
		dbStruct.DateOfBirth,
		dbStruct.Address,
		dbStruct.Department,
		dbStruct.AboutMe,
		dbStruct.Phone,
		dbStruct.Email,
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("row err: %v ", err)
		}
	}

	return id, err
}

func (db *dbClient) GetByID(ctx context.Context, id string) ([]byte, error) {
	rows, err := db.DB.QueryContext(ctx, "select id, firstname, lastname, middlename, date_of_birth,addres, department, about_me, phone, email from public.emploees where id in ($1)", id)
	if err != nil {
		return nil, err
	}

	var st models.DbStruct

	for rows.Next() {
		if err = rows.Scan(
			&st.ID,
			&st.FirstName,
			&st.LastName,
			&st.MiddleName,
			&st.DateOfBirth,
			&st.Address,
			&st.Department,
			&st.AboutMe,
			&st.Phone,
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
		return nil, fmt.Errorf("marshal error: %v", err)
	}

	return res, nil
}

func (db *dbClient) UpEmployee(ctx context.Context, id string, st *models.DbStruct) error {
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
