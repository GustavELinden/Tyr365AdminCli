package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

// requeueCmd represents the requeue command
var requeueCmd = &cobra.Command{
	Use:   "requeue",
	Short: "This command requeues a request in the Teams Governance API. Flag : --requestId number",
	Long: `This command requeues a request in the Teams Governance API. For example: 365Admin teamGov requeue --requestId 147999`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := Get("Requeue", map[string]string{"requestId": fmt.Sprintf("%d", requestId)})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	},
}

func init() {
	requeueCmd.Flags().Int32VarP(&requestId, "requestId", "r", 0, "The id of the request to requeue")
  if err := requeueCmd.MarkFlagRequired("requestId"); err != nil {
		fmt.Println(err)
	}

  TeamGovCmd.AddCommand(requeueCmd)

}
