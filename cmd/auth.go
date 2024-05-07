/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/skysyzygy/tq/auth"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var hostname, username, usergroup, location *string

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:     "authenticate",
	Aliases: []string{"a", "au", "aut", "auth"},
	Short:   "Authenticate with the Tessitura API",
	Long: `Manage authentication with various Tessitura API servers, 
	usernames and usergroups.`,
}

var authenticateAddCmd = &cobra.Command{
	Use:     "Add",
	Aliases: []string{"a", "add"},
	Short:   "Add an Tessitura API authentication method",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Password: ")
		var (
			password []byte
			err      error
		)
		inputFd := int(os.Stdin.Fd())
		if term.IsTerminal(inputFd) {
			password, err = term.ReadPassword(inputFd)
		} else {
			password, err = io.ReadAll(os.Stdin)
		}
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		// strip protocol from hostname if it exists
		host := regexp.MustCompile("^.+//").ReplaceAllString(*hostname, "")

		a := auth.New(host, *username, *usergroup, *location, password)
		err = a.Save()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	},
}

var authenticateListCmd = &cobra.Command{
	Use:     "authenticate",
	Aliases: []string{"a", "au", "aut", "auth"},
	Short:   "Authenticate with the Tessitura API",
	Long:    `List all saved Tessitura API authentication methods`,
	Run: func(cmd *cobra.Command, args []string) {
		if *hostname != "" ||
			*username != "" ||
			*usergroup != "" ||
			*location != "" {
			fmt.Println("Warning: parameters ignored")
		}
		auths, _ := auth.List()
		for _, auth := range auths {
			fmt.Printf("%s", auth)
		}
	},
}

var authenticateDeleteCmd = &cobra.Command{
	Use:     "authenticate",
	Aliases: []string{"a", "au", "aut", "auth"},
	Short:   "Authenticate with the Tessitura API",
	Long:    `Delete a Tessitura API authentication method`,
	Run: func(cmd *cobra.Command, args []string) {
		a := auth.New(*hostname, *username, *usergroup, *location, nil)
		a.Delete()
	},
}

func init() {

	rootCmd.AddCommand(authenticateCmd)

	hostname = authenticateCmd.PersistentFlags().StringP("host", "H", "", "hostname and base path of the API server")
	username = authenticateCmd.PersistentFlags().StringP("user", "u", "", "username to authenticate")
	usergroup = authenticateCmd.PersistentFlags().StringP("group", "g", "", "group to authenticate with")
	location = authenticateCmd.PersistentFlags().StringP("location", "l", "", "machine location to authenticate with")

	authenticateCmd.AddCommand(authenticateAddCmd, authenticateListCmd, authenticateDeleteCmd)
}
