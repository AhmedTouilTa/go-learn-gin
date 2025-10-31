package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("postgres",
		"user=gouser password=password dbname=gotransactions host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	Database = db
	return db
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return Database.Query(query, args...)
}

func Exec(query string, args ...any) (sql.Result, error) {
	return Database.Exec(query, args...)
}

func Ping() {
	err := Database.Ping()
	if err != nil {
		panic("Database not established !")
	}
}

func Close() {
	Database.Close()
}
