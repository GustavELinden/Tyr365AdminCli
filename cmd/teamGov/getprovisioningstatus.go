package teamGov

import (
	"fmt"

	teamGovHttp "github.com/GustavELinden/Tyr365AdminCli/TeamsGovernance"
	"github.com/spf13/cobra"
)

var getprovisioningstatusCmd = &cobra.Command{
	Use:   "getrequest",
	Short: "Get provisioning status of a request. Flag : --requestId number",
	Long:  `Get the provisioning status of a request. For example: 365Admin teamGov getprovisioningstatus --requestId 147999`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := teamGovHttp.Get("GetProvisioningStatus", map[string]string{"requestId": fmt.Sprintf("%d", requestId)})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		requests, err := teamGovHttp.UnmarshalRequests(&body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		ViewTable(&requests)
	},
}

func init() {

	getprovisioningstatusCmd.Flags().Int32VarP(&requestId, "requestId", "r", 0, "The id of the request to requeue")
	if err := getprovisioningstatusCmd.MarkFlagRequired("requestId"); err != nil {
		fmt.Println(err)
	}
	TeamGovCmd.AddCommand(getprovisioningstatusCmd)
}

// func RenderRequests(requests []teamGovHttp.Request) {
// 	// Create a table to display the response data
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader([]string{"ID", "Created", "GroupID", "TeamName", "Endpoint", "CallerID", "Status", "ProvisioningStep", "Message", "InitiatedBy", "Modified", "RetryCount", "QueuePriority"}) // Customize the table header as needed
// 	for _, req := range requests {
// 		row := []string{
// 			fmt.Sprintf("%d", req.ID),
// 			req.Created,
// 			req.GroupID,
// 			req.TeamName,
// 			req.Endpoint,
// 			req.CallerID,
// 			req.Status,
// 			req.ProvisioningStep,
// 			req.Message,
// 			req.InitiatedBy,
// 			req.Modified,
// 			fmt.Sprintf("%v", req.RetryCount),
// 			fmt.Sprintf("%d", req.QueuePriority),
// 		}
// 		table.Append(row)
// 	}
// 	table.Render()
// }
func ViewTable(d teamGovHttp.Printer) {
	d.PrintTable()
}