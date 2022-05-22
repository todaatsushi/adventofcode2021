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

var powerCmd = &cobra.Command{
	Use: "power",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		filename, err := flags.GetString("file")
		if filename == "" {
			log.Fatal("Couldn't get filename:", filename, err)
		}
		getPowerConsumption(filename)
	},
}

func getPowerConsumption(filename string) {
	allReadings := readings.ReadRawReadings(filename)
	gamma, epsilon := readings.GetGammaAndEpsilon(allReadings)
	powerConsumption := gamma * epsilon
	fmt.Println("Power consumption from readings:", powerConsumption)
}

func init() {
	rootCmd.AddCommand(powerCmd)
	powerCmd.Flags().StringP("file", "f", "", "Pass input file to get the power consumption.")
}
