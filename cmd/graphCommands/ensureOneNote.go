/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"
	"sync"

	models "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/spf13/cobra"
)

// ensureOneNoteCmd represents the ensureOneNote command
var ensureOneNoteCmd = &cobra.Command{
	Use:   "ensureOneNote",
	Short: "Calls all Onenotes in the team",
	Long:  `The command calls all channels in the team, and then calls each channels Onenote to ensure its initialized correctly`,
	Run: func(cmd *cobra.Command, args []string) {
		group, _ := graphHelper.GetGroupById(groupId)
		teamId := group.GetId()
		channels, _ := graphHelper.GetAllChannels(*teamId)
		allChannels := channels.GetValue()

		var wg sync.WaitGroup
		errors := make(chan error, len(allChannels))
		results := make(chan string, len(allChannels))

		for _, channel := range allChannels {
			wg.Add(1)
			go func(ch models.Channelable) {
				defer wg.Done()
				chanId := ch.GetId()
				chanTitle := ch.GetDisplayName()

				fmt.Println("Ensure files folders called for: " + *chanTitle)
				_, err := graphHelper.EnsureOneNote(*teamId, *chanId)
				if err != nil {
					errors <- fmt.Errorf("error calling %s: %v", *chanTitle, err)
					return
				}
				results <- "Successfully made call for " + *chanTitle
			}(channel)
		}

		wg.Wait()
		close(errors)
		close(results)

		for err := range errors {
			fmt.Println(err)
		}
		for res := range results {
			fmt.Println(res)
		}
	},
}

func init() {
	ensureOneNoteCmd.Flags().StringVarP(&groupId, "groupId", "r", "", "The id of the group")
	if err := ensureOneNoteCmd.MarkFlagRequired("groupId"); err != nil {
		fmt.Println(err)
	}
	GraphCmd.AddCommand(ensureOneNoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ensureOneNoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ensureOneNoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
