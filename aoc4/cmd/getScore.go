/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getScoreCmd = &cobra.Command{
	Use: "getScore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getScore called")
	},
}

func init() {
	rootCmd.AddCommand(getScoreCmd)
}
