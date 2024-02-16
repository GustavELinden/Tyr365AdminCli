/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getaverageprovisioningtimeCmd represents the getaverageprovisioningtime command
var getaverageprovisioningtimeCmd = &cobra.Command{
	Use:   "getaverageprovisioningtime",
	Short: "Get the average provisioning time in the Teams Governance API.",
	Long: `This command gets the average provisioning time in the Teams Governance API. For example: 365Admin teamGov getaverageprovisioningtime
	The time is an average of the time it takes from when a request is queued (Created time) to when it is succeeded (Modified time) `,
	Run: func(cmd *cobra.Command, args []string) {
				if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := Get("AverageProvisionTime")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		 fmt.Println("Raw response:", string(body))
		},
	}

func init() {
	TeamGovCmd.AddCommand(getaverageprovisioningtimeCmd)

}
