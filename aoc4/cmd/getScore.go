/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"internal/aoc4"
	"log"

	"github.com/spf13/cobra"
)

var getScoreCmd = &cobra.Command{
	Use: "getScore",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		filename, err := flags.GetString("file")
		if filename == "" {
			log.Fatal("Couldn't get filename:", filename, err)
		}
		aoc4.ReadInput(filename)
	},
}

func init() {
	rootCmd.AddCommand(getScoreCmd)
	getScoreCmd.Flags().StringP("file", "f", "", "Input file")
}
