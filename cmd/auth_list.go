/*
Copyright © 2024 Sky Syzygy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/99designs/keyring"
	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
	auth_table "github.com/ssyzygy/tq/cmd/auth/table"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "A brief description of your command",
	Long:    helpParagraph("Lists all authentication methods in the authentication storage"),
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {

		keys, _ := keyring.Open(keyring.Config{
			ServiceName: "tq",
		})

		keyList, _ := keys.Keys()
		rows := make([]table.Row, len(keyList))
		for i, key := range keyList {
			rows[i] = table.Row{key}
		}

		selected := auth_table.AuthTable(rows)

		fmt.Println(selected)

	},
}

func init() {
	authenticateCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}