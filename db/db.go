package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		log.Fatal("db url not set")
	}

	// pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// )

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	return db, nil
}
