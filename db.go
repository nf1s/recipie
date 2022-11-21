package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	os.Remove("sqlite.db")

	log.Println("Creating sqlite.db...")
	file, err := os.Create("sqlite.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite.db created")

	db, _ := sql.Open("sqlite3", "./sqlite.db")
	defer db.Close()

	createRecipesTable(db)
	createIngredientsTable(db)

	salmonAndPotatoes := Recipe{name: "Salmon and potatoes"}
	salmonAndPotatoes.insert(db)

	potatoes := Ingredient{name: "Potatoes", count: 4, recipeId: salmonAndPotatoes.id}
	potatoes.insert(db)

	Salmon := Ingredient{name: "Salmon", weightInGrams: 500, recipeId: salmonAndPotatoes.id}
	Salmon.insert(db)

	salmonAndPotatoes.display(db)
	potatoes.display(db)
	Salmon.display(db)

}
