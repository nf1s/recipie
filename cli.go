package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var CREATE_RECIPE = "Create a Recipe"
var LIST_RECIPIES = "List Recipes"
var EXIT = "exit"

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

func ingredient() int {
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
		return 0
	}
	fmt.Printf("name=%s, count=%d, weightInGrams=%d", answers.Name, answers.Count, answers.WeightInGrams)

	isAddMore := confirm("add more?")
	if isAddMore {
		ingredient()
	}

	return 1

}

func recipe() string {
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

func cli() int {
	choice := recipe()
	switch choice {
	case CREATE_RECIPE:
		fmt.Println("create")
		recipeName := input("what is the recipe's name?")
		fmt.Printf("recipe name: %s", recipeName)
		ingredient()
		return cli()
	case LIST_RECIPIES:
		fmt.Println("list")
		return cli()
	case EXIT:
		return 1
	default:
		return 1
	}
}
