package main

import (
	"database/sql"
	"log"
)

type Ingredient struct {
	name          string
	recipeId      string
	count         int
	weightInGrams int
}

var ingredientsTable string = `CREATE TABLE ingredients (
  "name" TEXT NOT NULL,
  "recipeId" integer NOT NULL,
  "count" integer,
  "weightInGrams" integer,
   PRIMARY KEY (name, recipeId),
   FOREIGN KEY (recipeId) REFERENCES recipes(id)
  );`

func createIngredientsTable(db *sql.DB) {
	log.Printf("Create ingredients table...")
	statement, err := db.Prepare(ingredientsTable) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("ingredients table created")

}

func (i Ingredient) create(db *sql.DB, name, recipeId string, count, weightInGrams int) *Ingredient {
	ing := &Ingredient{name: name, recipeId: recipeId, count: count, weightInGrams: weightInGrams}
	ing.insert(db)
	return ing
}

func (i *Ingredient) insert(db *sql.DB) {
	log.Println("Inserting ingredient record ...")
	insertStudentSQL := `INSERT INTO ingredients (name, recipeId, count, weightInGrams)  VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(i.name, i.recipeId, i.count, i.weightInGrams)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (i *Ingredient) display(db *sql.DB) {
	row, err := db.Query("SELECT * FROM ingredients Where name=?", i.name)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&i.name, &i.recipeId)
		log.Printf("Ingredient: %s %s %d %d", i.name, i.recipeId, i.count, i.weightInGrams)
	}
}
