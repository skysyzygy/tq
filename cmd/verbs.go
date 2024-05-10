package cmd

import (
	"github.com/spf13/cobra"
)

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
	rootCmd.AddCommand(Get_cmd, Post_cmd, Put_cmd)
}
