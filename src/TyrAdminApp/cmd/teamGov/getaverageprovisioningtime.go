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
	Use:   "getavg",
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


//https://teams.microsoft.com/l/team/19%3AsfBeq4-sPDdPSr8rt-EKBH4ee4SMwAETyMOdQX0iXnc1%40thread.tacv2/conversations?groupId=3e763cb7-606f-4aaa-a6c2-693833c0f1ca&tenantId=fd6c9945-d5d1-423c-acf9-2a5431314398