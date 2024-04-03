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
	"strings"

	"github.com/99designs/keyring"
	"github.com/spf13/cobra"
	auth_textinput "github.com/ssyzygy/tq/cmd/auth/textinput"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an authentication method to the authentication storage",
	Long:  helpParagraph("Adds an authentication method to the authentication storage"),

	Run: func(cmd *cobra.Command, args []string) {

		keys, _ := keyring.Open(keyring.Config{
			ServiceName: "tq",
		})

		values, err := auth_textinput.AuthInput()
		if err != nil {
			fmt.Println(hiliteStyle.Render(fmt.Sprintf("Error: %s", err.Error())))
			return
		}

		if err := keys.Set(keyring.Item{
			Key:  strings.Join(values[0:4], ":"),
			Data: []byte(values[4]),
		}); err != nil {
			fmt.Println(err)
		}

		key, _ := keys.Get(strings.Join(values[0:4], ":"))

		fmt.Println(successStyle.Render("Added key to the keyring:"))
		fmt.Println(" * " + key.Key)
		fmt.Println(" * " + string(key.Data))

	},
}

func init() {
	authenticateCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
