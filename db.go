package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	os.Remove("sqlite.db")

	log.Println("Creating sqlite.db...")
	file, err := os.Create("sqlite.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite.db created")

	db, _ := sql.Open("sqlite3", "./sqlite.db")
	return db

}
