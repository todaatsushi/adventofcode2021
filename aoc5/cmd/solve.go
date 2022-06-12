package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use: "solve",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solve called")
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
}
