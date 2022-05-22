/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"
	"internal/readings"
	"log"

	"github.com/spf13/cobra"
)

var lifeSupportCmd = &cobra.Command{
	Use: "lifeSupport",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		filename, err := flags.GetString("file")
		if filename == "" {
			log.Fatal("Couldn't get filename:", filename, err)
		}
		getLifeSupportRating(filename)
	},
}

func getLifeSupportRating(filename string) {
	allReadings := readings.ReadRawReadings(filename)
	fmt.Println(allReadings)
}

func init() {
	rootCmd.AddCommand(lifeSupportCmd)
	lifeSupportCmd.Flags().StringP("file", "f", "", "Pass input file to get the life support rating.")
}
