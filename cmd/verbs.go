package cmd

import (
	"github.com/spf13/cobra"
)

var Get_cmd = &cobra.Command{
	Use:     "Get",
	Aliases: []string{"Get", "get", "G", "g"},
	Short:   "Retrieve entities from Tessitura",
}

var Post_cmd = &cobra.Command{
	Use:     "Post",
	Aliases: []string{"Post", "post", "P", "p"},
	Short:   "Create entities in Tessitura",
}

var Put_cmd = &cobra.Command{
	Use:     "Put",
	Aliases: []string{"put", "P", "p", "Put"},
	Short:   "Update entities in Tessitura",
}

func init() {
	rootCmd.AddCommand(Get_cmd, Post_cmd, Put_cmd)
}
