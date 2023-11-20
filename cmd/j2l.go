/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/agoulart72/leveldb-convert/convert"
	"github.com/spf13/cobra"
)

// j2lCmd represents the j2l command
var j2lCmd = &cobra.Command{
	Use:   "j2l",
	Short: "Convert from json to leveldb",
	Long:  `Convert from json to leveldb.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return convert.ConvertJsonToLeveldb(args)
	},
}

func init() {
	rootCmd.AddCommand(j2lCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// j2lCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// j2lCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
