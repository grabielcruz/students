package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	godotenv.Load(".env")
	DBUSER := os.Getenv("DBUSER")
	PASSWORD := os.Getenv("PASSWORD")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")

	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v", DBUSER, PASSWORD, HOST, PORT, DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
