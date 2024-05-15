package cmd

import (
	"github.com/spf13/cobra"
)

const longSuffix string = `Queries are simply JSON objects and can be batched by ` +
	`combining multiple query objects into a single JSON array, e.g.\n` +
	`[{"ID":123},{"ID":124},{"ID":125},...]` +
	`Object definitions are detailed in the help for each command.`

var Get_cmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"Get", "get", "Retrieve", "retrieve", "G", "g", "R", "r"},
	Short:   "Retrieve entities from Tessitura",
}

var Post_cmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"Post", "post", "Create", "create", "C", "c"},
	Short:   "Create entities in Tessitura",
}

var Put_cmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"Put", "put", "Update", "update", "U", "u"},
	Short:   "Update entities in Tessitura",
}

func init() {
	Get_cmd.Long = Get_cmd.Short + "\n" + longSuffix
	Post_cmd.Long = Post_cmd.Short + "\n" + longSuffix
	Put_cmd.Long = Put_cmd.Short + "\n" + longSuffix
	rootCmd.AddCommand(Get_cmd, Post_cmd, Put_cmd)
}
