/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:     "authenticate",
	Aliases: []string{"a", "au", "aut", "auth"},
	Short:   "Authenticate with the Tessitura API",
	Long: helpParagraph(`A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("authenticate called")
	},
}

func init() {

	rootCmd.AddCommand(authenticateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authenticateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authenticateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
