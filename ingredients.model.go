package main

import (
	"database/sql"
	"log"
)

var ingredientsTable string = `CREATE TABLE ingredients (
  "name" TEXT NOT NULL,
  "recipeId" integer NOT NULL,
  "count" integer,
  "weightInGrams" integer,
   PRIMARY KEY (name, recipeId),
   FOREIGN KEY (recipeId) REFERENCES recipes(id)
  );`

type Ingredient struct {
	name          string
	recipeId      string
	count         int
	weightInGrams int
}

type IngredientRepository interface {
	CreateIngredientsTable()
	Get(name string) *sql.Row
	Insert(name string, recipeId string, count int, weightInGrams int) *Ingredient
}

type ingredientRepository struct {
	DB *sql.DB
}

func NewIngredientRepository(db *sql.DB) IngredientRepository {
	return &ingredientRepository{DB: db}
}
func (i *ingredientRepository) CreateIngredientsTable() {
	log.Printf("Create ingredients table...")
	statement, err := i.DB.Prepare(ingredientsTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("ingredients table created")

}

func (i *ingredientRepository) Insert(name string, recipeId string, count int, weightInGrams int) *Ingredient {
	ing := Ingredient{name: name, recipeId: recipeId, count: count, weightInGrams: weightInGrams}
	log.Println("Inserting ingredient record ...")
	insertStudentSQL := `INSERT INTO ingredients (name, recipeId, count, weightInGrams)  VALUES (?, ?, ?, ?)`
	statement, err := i.DB.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(&ing.name, &ing.recipeId, &ing.count, &ing.weightInGrams)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &ing
}

func (i *ingredientRepository) Get(name string) *sql.Row {
	var ing Ingredient
	row := i.DB.QueryRow("SELECT * FROM ingredients Where name=?", ing.name)
	err := row.Scan(&ing.name, &ing.recipeId)
	if err != nil {
		log.Fatal(err)
	}
	return row

}
