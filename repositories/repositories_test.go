package repositories

import (
	"context"
	"fmt"
	"github.com/AlexandrSminrov/employees/models"
	"github.com/joho/godotenv"
	"os"
	"testing"
	"time"
)

func InitDB() error {

	if err := godotenv.Load("../.env"); err != nil {
		return fmt.Errorf("No .env file found ")
	}

	if err := os.Setenv("pgHost", "localhost"); err != nil {
		return err
	}

	if err := ConnectDB(); err != nil {
		return fmt.Errorf("Error ConnectDB: %v ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := ConnDb.DB.PingContext(ctx); err != nil {
		return fmt.Errorf("Base not started!!! ")
	}
	return nil
}

func TestDbQuery_GetAll(t *testing.T) {
	if err := InitDB(); err != nil {
		t.Fatal(err)
	}
	defer func(t *testing.T) {
		if err := ConnDb.DB.Close(); err != nil {
			t.Fatalf("Conn DB clode error: %v ", err)
		}
	}(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	st, err := ConnDb.GetAll(ctx)
	if err != nil {
		t.Error(err)
	}

	for i, r := range st {
		if len(r.AboutMe) > 1 || len(r.Address) > 1 {
			t.Errorf("Struct %d error this method should not return fields\n\tAboutMe Len: %d\tAboutMe: %s\n\t\tAddress len: %d\t Address: %s",
				i,
				len(r.AboutMe),
				r.AboutMe,
				len(r.Address),
				r.Address,
			)
		}
	}

}

func TestDbQuery_AddEmployee(t *testing.T) {
	if err := InitDB(); err != nil {
		t.Fatal(err)
	}
	defer func(t *testing.T) {
		if err := ConnDb.DB.Close(); err != nil {
			t.Fatalf("Conn DB clode error: %v ", err)
		}
	}(t)

	st := models.DbStruct{
		FirstName:  "Иванов",
		LastName:   "Иван",
		BDate:      "12.31.1922",
		Address:    "Москва",
		Department: "HR",
		AboutMe:    "я",
		Tnumber:    "71234567890",
		Email:      "exaple@ex.ex",
	}
	if err := st.Validate(); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id, err := ConnDb.AddEmployee(&st, ctx)

	if id < 1 {
		t.Errorf("The function should return id\tid: %d ", id)
	}

	if err != nil {
		t.Error(err)
	}
}
