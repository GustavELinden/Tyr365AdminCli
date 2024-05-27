/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"encoding/json"
	"fmt"
	"sync"

	saveToFile "github.com/GustavELinden/Tyr365AdminCli/SaveToFile"
	teamGovHttp "github.com/GustavELinden/Tyr365AdminCli/TeamsGovernance"
	"github.com/spf13/cobra"
)

type ApiResponse struct {
	StatusCode int `json:"statusCode"`
}

var fileName string

// readFileCmd represents the readFile command
var readFileCmd = &cobra.Command{
	Use:   "readFile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var readGroups interface{}
		saveToFile.ReadDataFromJSONFile(fileName, &readGroups)

		if cmd.Flag("output").Changed {
			outData, _ := json.Marshal(readGroups)
			fmt.Println(string(outData))
		}
		if cmd.Flag("updateCT").Changed {
			updateContenTypes(readGroups)
		}
	},
}

func init() {
	readFileCmd.Flags().StringVarP(&fileName, "file", "f", "", "The name of the file you want to read from")
	readFileCmd.Flags().Bool("updateCT", false, "Call updateCTs in TeamGOv API")
	TeamGovCmd.AddCommand(readFileCmd)
}

func updateContenTypes(readGroups interface{}) {
	outData, _ := json.Marshal(readGroups)
	requests, err := teamGovHttp.UnmarshalRequests(&outData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a channel to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, 5)

	var wg sync.WaitGroup
	for _, group := range requests {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a token

		// Start a new goroutine for each group
		go func(group teamGovHttp.Request) {
			defer wg.Done()
			defer func() { <-semaphore }()

			queryParams := make(map[string]string)
			queryParams["groupId"] = group.GroupID
			_, err := teamGovHttp.Get("SetContentTypesToEditOnSite", queryParams)

			if err != nil {
				fmt.Printf("Failed to process group %s: %v\n", group.GroupID, err)
				return
			}

			fmt.Printf("Successfully processed group %s\n", group.GroupID)

		}(group)
	}

	// Wait for all goroutines to complete
	wg.Wait()

}
