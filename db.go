package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createDB() {

	log.Println("Creating sqlite.db...")
	file, err := os.Create("sqlite.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite.db created")

}

func InitDB() *sql.DB {

	db, err := sql.Open("sqlite3", "./sqlite.db")

	if os.IsNotExist(err) {
		createDB()
	}

	return db

}
