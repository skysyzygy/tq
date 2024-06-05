package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

const longSuffix string = "Queries are simply JSON objects and can be batched by " +
	"combining multiple query objects into a single JSON array, e.g. \n\n\t" +
	`[{"ID":123},{"ID":124},{"ID":125},...]` +
	"\n\nQuery details are detailed in the help for each command."

var Get_cmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"get", "retrieve", "g", "r"},
	Short:   "Retrieve entities from Tessitura",
}

var Post_cmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"post", "create", "c"},
	Short:   "Create entities in Tessitura",
}

var Put_cmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"put", "update", "u"},
	Short:   "Update entities in Tessitura",
}

func init() {
	Get_cmd.Long = Get_cmd.Short + "\n" + longSuffix
	Post_cmd.Long = Post_cmd.Short + "\n" + longSuffix
	Put_cmd.Long = Put_cmd.Short + "\n" + longSuffix

	usageTemplate := strings.NewReplacer("verb", "object", "Verb", "Object").
		Replace(rootCmd.UsageTemplate())
	Get_cmd.SetUsageTemplate(usageTemplate)
	Post_cmd.SetUsageTemplate(usageTemplate)
	Put_cmd.SetUsageTemplate(usageTemplate)

	rootCmd.AddCommand(Get_cmd, Post_cmd, Put_cmd)

}
