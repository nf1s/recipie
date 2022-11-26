package cli

import "github.com/AlecAivazis/survey/v2"

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
