package cmd

import (
	"internal/aoc5"
	"log"

	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use: "solve",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		filename, err := flags.GetString("file")
		if filename == "" || err != nil {
			log.Fatal("Couldn't get filename:", filename, err)
		}
		aoc5.Solve(filename)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
	solveCmd.Flags().StringP("file", "f", "", "Input file")
}
