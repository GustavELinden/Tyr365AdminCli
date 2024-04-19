// }
package graphCommands

import (
	"fmt"
	"sync"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/spf13/cobra"
)

var groupId string

// ensureFilesFolderCmd represents the ensureFilesFolder command
var ensureFilesFolderCmd = &cobra.Command{
	Use:   "ensureFilesFolder",
	Short: "Ensure Files Folder is present in all channels of a given team",
	Long:  `This command checks and ensures that a Files Folder is present in all channels of a specified team in Microsoft Teams.`,
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
				_, err := graphHelper.EnsureFilesFolder(*teamId, *chanId)
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
	ensureFilesFolderCmd.Flags().StringVarP(&groupId, "groupId", "r", "", "The id of the group")
	if err := ensureFilesFolderCmd.MarkFlagRequired("groupId"); err != nil {
		fmt.Println(err)
	}
	GraphCmd.AddCommand(ensureFilesFolderCmd)
}
