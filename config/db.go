package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func ConnectDB() *pg.DB {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing required environment variables")
	}

	db := pg.Connect(&pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	// defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS metrics (
		id SERIAL PRIMARY KEY,
		timestamp BIGINT NOT NULL,
		cpu_load FLOAT NOT NULL,
		concurrency INTEGER NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully")

	return db
}
