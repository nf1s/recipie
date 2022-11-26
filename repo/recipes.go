package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

var recipesTable string = `CREATE TABLE IF NOT EXISTS recipes(
  "id" TEXT NOT NULL PRIMARY KEY,
  "name" TEXT
  );`

type Recipe struct {
	Id   string
	Name string
}

type RecipeRepository interface {
	CreateRecipesTable()
	Insert(name string) *Recipe
	ListRecipes() *sql.Rows
	Get(id string) *sql.Row
	GetFullRecipe(id string) *sql.Row
}

type recipeRepository struct {
	DB *sql.DB
}

func NewRecipeRepository(db *sql.DB) RecipeRepository {
	return &recipeRepository{DB: db}
}

func (r *recipeRepository) CreateRecipesTable() {
	log.Printf("Create recipes table...")
	statement, err := r.DB.Prepare(recipesTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Recipes table created")

}

func (r *recipeRepository) ListRecipes() *sql.Rows {
	var recipe Recipe
	row, err := r.DB.Query("SELECT * FROM recipes")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&recipe.Id, &recipe.Name)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}
	return row

}

func (r *recipeRepository) Insert(name string) *Recipe {
	recipe := Recipe{Id: fmt.Sprintf("%s", uuid.New()), Name: name}
	log.Println("Inserting recipe record ...")
	insertStudentSQL := `INSERT INTO recipes (id, name)  VALUES (?,?)`
	statement, err := r.DB.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(&recipe.Id, &recipe.Name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &recipe
}

func (r *recipeRepository) Get(id string) *sql.Row {
	var recipe Recipe
	row := r.DB.QueryRow("SELECT * FROM recipes WHERE id=?", id)
	err := row.Scan(&recipe.Id, &recipe.Name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return row

}

func (r *recipeRepository) GetFullRecipe(id string) *sql.Row {
	var recipe Recipe
	row := r.DB.QueryRow("SELECT * FROM recipes JOIN ingredients on recipes.id=ingredients.recipeId WHERE recipes.id=?", id)
	err := row.Scan(&recipe.Id, &recipe.Name)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return row

}
