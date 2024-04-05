package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

// requeueCmd represents the requeue command
var retryRequestCmd = &cobra.Command{
	Use:   "retryRequest",
	Short: "This command requeues a request in the Teams Governance API. Flag : --requestId number",
	Long:  `This command requeues a request in the Teams Governance API. For example: 365Admin teamGov retryRequest --requestId 147999`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := Post("RetryRequest", map[string]string{"requestId": fmt.Sprintf("%d", requestId)})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	},
}

func init() {
	retryRequestCmd.Flags().Int32VarP(&requestId, "requestId", "r", 0, "The id of the request to requeue")
	if err := retryRequestCmd.MarkFlagRequired("requestId"); err != nil {
		fmt.Println(err)
	}

	TeamGovCmd.AddCommand(retryRequestCmd)

}
