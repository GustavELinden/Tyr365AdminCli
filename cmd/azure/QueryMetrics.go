/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package azure

import (
	"fmt"

	azurehelper "github.com/GustavELinden/Tyr365AdminCli/AzureHelper"
	"github.com/spf13/cobra"
)

// QueryMetricsCmd represents the QueryMetrics command
var QueryMetricsCmd = &cobra.Command{
	Use:   "QueryMetrics",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		responses, err := azurehelper.GetMetrics()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*responses)
	},
}

func init() {
	AzureCmd.AddCommand(QueryMetricsCmd)
	// Define flags and configuration settings as needed
}
