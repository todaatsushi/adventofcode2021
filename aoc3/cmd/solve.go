/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"internal/readings"
	"log"

	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "A brief description of your command",
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
	rootCmd.AddCommand(solveCmd)

	flags := solveCmd.Flags()
	flags.StringP("file", "f", "", "Pass input file to get the power consumption.")
}
