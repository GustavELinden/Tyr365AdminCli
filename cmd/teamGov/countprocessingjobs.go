/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

// countprocessingjobsCmd represents the countprocessingjobs command
var countprocessingjobsCmd = &cobra.Command{
	Use:   "countp",
	Short: "The commands counts the number of processing jobs in the Teams Governance API.",
	Long:  `This command counts the number of processing jobs in the Teams Governance API. For example: 365Admin teamGov countprocessingjobs`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := Get("GetProcessingRequests")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		requests, err := UnmarshalInteger(&body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(requests)
	},
}

func init() {
	TeamGovCmd.AddCommand(countprocessingjobsCmd)
}
