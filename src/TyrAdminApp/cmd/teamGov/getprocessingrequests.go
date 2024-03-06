package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getprocessingrequestsCmd represents the getprocessingrequests command
var getprocessingrequestsCmd = &cobra.Command{
	Use:   "countprequests",
	Short: "Get the number of processing requests in the Teams Governance API.",
	Long: `This command gets the number of processing requests in the Teams Governance API. For example: 365Admin teamGov getprocessingrequests`,
	Run: func(cmd *cobra.Command, args []string) {
		//If flag --help is used, print the help message
		if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := Get("GetProcessingRequests")
if err != nil {
	fmt.Println("Error:", err)
	return
}
requests, err := UnmarshalInteger(&body);
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println(requests)
	},
}

func init() {
	TeamGovCmd.AddCommand(getprocessingrequestsCmd)

}
