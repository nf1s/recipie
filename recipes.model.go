package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Recipe struct {
	id   string
	name string
}

var recipesTable string = `CREATE TABLE recipes (
  "id" TEXT NOT NULL PRIMARY KEY,
  "name" TEXT
  );`

func createRecipesTable(db *sql.DB) {
	log.Printf("Create recipes table...")
	statement, err := db.Prepare(recipesTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Recipes table created")

}

func (r *Recipe) insert(db *sql.DB) *Recipe {
	r.id = fmt.Sprintf("%s", uuid.New())
	log.Println("Inserting recipe record ...")
	insertStudentSQL := `INSERT INTO recipes (id, name)  VALUES (?,?)`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(r.id, r.name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return r
}

func (r *Recipe) display(db *sql.DB) {
	row, err := db.Query("SELECT * FROM recipes WHERE id=?", r.id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&r.id, &r.name)
		log.Printf("Recipe: %s %s", r.id, r.name)

	}
}

func (r *Recipe) getFullRecipe(db *sql.DB) {
	row, err := db.Query("SELECT * FROM recipes JOIN ingredients on recipes.id=ingredients.recipeId WHERE recipes.id=?", r.id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&r.id, &r.name)
		log.Printf("Recipe: %s %s", r.id, r.name)

	}
}
