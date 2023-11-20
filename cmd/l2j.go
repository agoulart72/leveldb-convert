/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/agoulart72/leveldb-convert/convert"
	"github.com/spf13/cobra"
)

// l2jCmd represents the l2j command
var l2jCmd = &cobra.Command{
	Use:   "l2j",
	Short: "Convert from leveldb to json",
	Long:  `Convert from leveldb to json`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return convert.ConvertLeveldbToJson(args)
	},
}

func init() {
	rootCmd.AddCommand(l2jCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// l2jCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// l2jCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
