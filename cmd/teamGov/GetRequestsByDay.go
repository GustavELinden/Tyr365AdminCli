/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	teamGovHttp "github.com/GustavELinden/Tyr365AdminCli/TeamsGovernance"
	"github.com/spf13/cobra"
)

var date string

// GetRequestsByDayCmd represents the GetRequestsByDay command
var GetRequestsByDayCmd = &cobra.Command{
	Use:   "GetRequestsByDay",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)
		if date != "" {
			queryParams["date"] = date
		}

		body, err := teamGovHttp.Get("GetRequestsForDay", queryParams)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		requests, err := teamGovHttp.UnmarshalRequests(&body)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		renderRequests(requests)

	},
}

func init() {
	GetRequestsByDayCmd.Flags().StringVarP(&date, "date", "d", "", "Date in fashion YYYY/MM/DD")
	TeamGovCmd.AddCommand(GetRequestsByDayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetRequestsByDayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetRequestsByDayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
