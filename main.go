package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

var rootCmd = &cobra.Command{
	Use:              "recipie [OPTIONS] [COMMANDS]",
	TraverseChildren: true,
	Short:            "Recipe CLI to handle multiple recipies and generate shopping list for a week or two",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Recipie",
	Long:  `All software has versions. This is Recipie's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Recipe v0.1.0 %s", args[0])
	},
}

var ingredientCmd = &cobra.Command{
	Use:   "ingredient",
	Short: "manage ingredients",
	Long:  `add, modify or remove an ingredient`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Ingredient %s", args[0])
	},
}

func Init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(ingredientCmd)
	ingredientCmd.PersistentFlags().StringP("name", "n", "Ingredient name", "")
	ingredientCmd.MarkFlagRequired("name")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
