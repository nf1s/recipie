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

	r := NewRecipeRepository(db)
	r.CreateRecipesTable()
	i := NewIngredientRepository(db)
	i.CreateIngredientsTable()

	salmonAndPotatoes := r.Insert("Salmon and Potatoes")

	potatoes := i.Insert("Potatoes", salmonAndPotatoes.id, 10, 0)

	salmon := i.Insert("Salmon", salmonAndPotatoes.id, 0, 500)

	r.Get(salmonAndPotatoes.id)
	r.Get(potatoes.name)
	r.Get(salmon.name)
}
