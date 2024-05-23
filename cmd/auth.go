/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/skysyzygy/tq/auth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

func init() {
	if term.IsTerminal(int(os.Stdin.Fd())) {
		pr = termPasswordReader{}
	} else {
		pr = stdinPasswordReader{}
	}
}

var hostname, username, usergroup, location *string
var pr passwordReader

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:     "authenticate",
	Aliases: []string{"a", "auth"},
	Short:   "Authenticate with the Tessitura API",
	Long:    `Manage authentication with various Tessitura API servers, usernames and usergroups.`,
}

type passwordReader interface {
	ReadPassword() ([]byte, error)
}

type (
	termPasswordReader  struct{}
	stdinPasswordReader struct{}
)

func (pr termPasswordReader) ReadPassword() ([]byte, error) {
	return term.ReadPassword(int(os.Stdin.Fd()))
}

func (pr stdinPasswordReader) ReadPassword() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}

var authenticateAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "add"},
	Short:   "Add an Tessitura API authentication method",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Password: ")
		var (
			password []byte
			err      error
		)
		password, err = pr.ReadPassword()
		if err != nil {
			return err
		}

		a := auth.New(*hostname, *username, *usergroup, *location, password)
		err = a.Save()
		return err
	},
}

var authenticateListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   `List all saved Tessitura API authentication methods`,
	Run: func(cmd *cobra.Command, args []string) {
		if *hostname != "" ||
			*username != "" ||
			*usergroup != "" ||
			*location != "" {
			os.Stderr.WriteString("Warning: parameters ignored\n")
		}
		auths, _ := auth.List()
		for _, auth := range auths {
			fmt.Printf("%s", auth)
		}
	},
}

var authenticateDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "del", "rm"},
	Short:   `Delete a Tessitura API authentication method`,
	RunE: func(cmd *cobra.Command, args []string) error {
		a := auth.New(*hostname, *username, *usergroup, *location, nil)
		return a.Delete()
	},
}

var authenticateSelectCmd = &cobra.Command{
	Use:     "select",
	Aliases: []string{"s", "sel"},
	Short:   `Select a Tessitura API authentication method`,
	RunE: func(cmd *cobra.Command, args []string) error {
		a := auth.New(*hostname, *username, *usergroup, *location, nil)
		err := a.Load()
		if err != nil {
			return err
		}
		str, _ := a.String()
		viper.Set("login", str)
		err = viper.WriteConfig()
		return err
	},
}

var authenticateValidateCmd = &cobra.Command{
	Use:     "validate",
	Aliases: []string{"v", "val"},
	Short:   `Validate a Tessitura API authentication method with the server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		a := auth.New(*hostname, *username, *usergroup, *location, nil)
		err := a.Load()
		if err != nil {
			return err
		}
		valid, err := a.Validate()
		if valid {
			os.Stderr.WriteString("Success: authentication is valid!")
		} else {
			os.Stderr.WriteString("Failure: authentication is not valid.")
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(authenticateCmd)

	hostname = authenticateCmd.PersistentFlags().StringP("host", "H", "", "hostname and base path of the API server")
	username = authenticateCmd.PersistentFlags().StringP("user", "u", "", "username to authenticate")
	usergroup = authenticateCmd.PersistentFlags().StringP("group", "g", "", "group to authenticate with")
	location = authenticateCmd.PersistentFlags().StringP("location", "L", "", "machine location to authenticate with")

	authenticateCmd.AddCommand(authenticateAddCmd, authenticateListCmd,
		authenticateDeleteCmd, authenticateSelectCmd, authenticateValidateCmd)
}
