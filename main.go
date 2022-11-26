package main

import (
	cli "recipie/cli"
	repo "recipie/repo"
)

func main() {
	db := InitDB()
	defer db.Close()

	recipe := repo.NewRecipeRepository(db)
	ingredient := repo.NewIngredientRepository(db)

	recipe.CreateRecipesTable()
	ingredient.CreateIngredientsTable()

	cli := cli.CLI{Recipe: recipe, Ingredient: ingredient}
	cli.Start()
}
