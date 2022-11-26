package main

import (
	"fmt"

	repo "recipie/repo"

	"github.com/AlecAivazis/survey/v2"
)

var CREATE_RECIPE = "Create a Recipe"
var LIST_RECIPIES = "List Recipes"
var EXIT = "exit"

type CLI struct {
	Recipe     repo.RecipeRepository
	Ingredient repo.IngredientRepository
}

func confirm(message string) bool {
	condition := false
	prompt := &survey.Confirm{
		Message: message,
	}
	survey.AskOne(prompt, &condition)
	return condition

}
func input(message string) string {
	var name string
	prompt := &survey.Input{
		Message: message,
	}
	survey.AskOne(prompt, &name)
	return name
}

func (cli *CLI) addIngredients(recipeId string) bool {
	var qs = []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "What is the name of the ingredient?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name:   "count",
			Prompt: &survey.Input{Message: "how many do you need?"},
		},
		{
			Name:   "weight in grams",
			Prompt: &survey.Input{Message: "how much do you need?"},
		},
	}
	fmt.Print("here")
	answers := struct {
		Name          string
		Count         int
		WeightInGrams int `survey:"weight in grams"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Printf("name=%s, count=%d, weightInGrams=%d", answers.Name, answers.Count, answers.WeightInGrams)
	cli.Ingredient.Insert(answers.Name, recipeId, answers.Count, answers.WeightInGrams)

	isAddMore := confirm("add more?")
	if isAddMore {
		cli.addIngredients(recipeId)
	}

	return true

}

func (cli *CLI) mainOptions() string {
	choice := ""
	prompt := &survey.Select{
		Message: "Choose an option:",
		Options: []string{CREATE_RECIPE, LIST_RECIPIES, EXIT},
		Default: LIST_RECIPIES,
	}
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return choice
}

func (cli *CLI) createRecipe() *repo.Recipe {
	fmt.Println("create a new recipie")
	recipeName := input("what is the recipe's name?")
	recipe := cli.Recipe.Insert(recipeName)
	fmt.Printf("recipe name: %s", recipeName)
	return recipe
}

func (cli *CLI) Start() bool {
	choice := cli.mainOptions()
	switch choice {
	case CREATE_RECIPE:
		recipe := cli.createRecipe()
		cli.addIngredients(recipe.Id)
		return cli.Start()
	case LIST_RECIPIES:
		fmt.Println("list")
		return cli.Start()
	case EXIT:
		return true
	default:
		return true
	}
}
