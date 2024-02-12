/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	getgov "github.com/GustavELinden/TyrAdminCli/365Admin/httpFuncs"
	"github.com/spf13/cobra"
)

// getaverageprovisioningtimeCmd represents the getaverageprovisioningtime command
var getaverageprovisioningtimeCmd = &cobra.Command{
	Use:   "getaverageprovisioningtime",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getaverageprovisioningtime called")
	
		body, err := getgov.Get("AverageProvisionTime")
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
