package cmd

import (
	"adventofcode2021/aoc5/internal/aoc5"
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

		val, err := flags.GetInt("value")
		if err != nil {
			log.Fatal("Couldn't get value:", err)
		}

		withDiagonals, err := flags.GetBool("value")
		if err != nil {
			log.Fatal("Couldn't get flag:", err)
		}

		aoc5.Solve(filename, val, withDiagonals)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
	solveCmd.Flags().StringP("file", "f", "", "Input file")
	solveCmd.Flags().IntP("value", "v", 420, "Which value to check the tally of")
	solveCmd.Flags().BoolP("with-diagonals", "d", true, "Count diagonal lines in totals")
}
