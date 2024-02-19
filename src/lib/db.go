package lib

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func conn() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error take db in .env file")
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_CONN"))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

var DB *sqlx.DB = conn()
