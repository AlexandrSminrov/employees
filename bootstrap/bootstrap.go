package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("No .env file found ")
	}
	return nil
}
