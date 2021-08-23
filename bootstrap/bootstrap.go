package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("No .env file found")
	}
	return nil
}

//func InitConnectDB() error {
//	dbStr := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		"localhost", 5432,
//		os.Getenv("pgUser"),
//		os.Getenv("pgPass"),
//		os.Getenv("pgDb"),
//	)
//
//	db, err := sql.Open("postgres", dbStr)
//	if err != nil {
//		return fmt.Errorf("Connection error: %v ", err)
//	}
//
//	maxConns, err := strconv.Atoi(os.Getenv("MaxConns"))
//	if err != nil {
//		return fmt.Errorf("MaxConns read error: %v ", err)
//	}
//
//	idleConns, err := strconv.Atoi(os.Getenv("IdleConns"))
//	if err != nil {
//		return fmt.Errorf("IdleConns read error: %v ", err)
//	}
//
//	db.SetMaxOpenConns(maxConns)
//	db.SetMaxIdleConns(idleConns)
//
//	repositories.ConnDb = db
//
//	return nil
//}
