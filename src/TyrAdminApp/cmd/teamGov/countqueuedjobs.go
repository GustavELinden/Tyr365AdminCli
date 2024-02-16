/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

// countqueuedjobsCmd represents the countqueuedjobs command
var countqueuedjobsCmd = &cobra.Command{
	Use:   "countqueuedjobs",
	Short: "counts how many jobs are queued in the Teams Governance API.",
	Long: `This command counts how many jobs are queued in the Teams Governance API. For example: 365Admin teamGov countqueuedjobs`,
	Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := Get("CountQueuedJobs")
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
  resp, err := UnmarshalInteger(&body);
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
  fmt.Println("Number of queued jobs: %s ", resp)
	},
}

func init() {
	TeamGovCmd.AddCommand(countqueuedjobsCmd)
}
