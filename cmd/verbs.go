package cmd

import (
	"github.com/spf13/cobra"
)

var Get_cmd = &cobra.Command{
	Use:     "Get",
	Aliases: []string{"Get", "get", "Retrieve", "retrieve", "R", "r"},
	Short:   "Retrieve entities from Tessitura",
}

var Post_cmd = &cobra.Command{
	Use:     "Post",
	Aliases: []string{"Post", "post", "Create", "create", "C", "c"},
	Short:   "Create entities in Tessitura",
}

var Put_cmd = &cobra.Command{
	Use:     "Put",
	Aliases: []string{"Put", "put", "Update", "update", "U", "u"},
	Short:   "Update entities in Tessitura",
}

func init() {
	rootCmd.AddCommand(Get_cmd, Post_cmd, Put_cmd)
}
