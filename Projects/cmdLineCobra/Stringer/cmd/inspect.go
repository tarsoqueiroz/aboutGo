/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"cmdlnCobraStringer/pkg/inspect"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp", "verify"},
	Short:   "Inspects a string",
	Long: `Inspects 
a string and show information about it.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := inspect.Inspect(i, false)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}

		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
