package teamGov

import (
	"fmt"
	"os"

	getgov "github.com/GustavELinden/TyrAdminCli/365Admin/httpFuncs"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// getprovisioningstatusCmd represents the getprovisioningstatus command


var getprovisioningstatusCmd = &cobra.Command{
    Use:   "getprovisioningstatus",
    Short: "Get provisioning status of a request. Flag : --requestId number",
    Long:  `Get the provisioning status of a request. For example: 365Admin teamGov getprovisioningstatus --requestId 147999`,
    Run: func(cmd *cobra.Command, args []string) {
        // Convert int32 requestId to string before passing it to the Get function
body, err := getgov.Get("GetProvisioningStatus", map[string]string{"requestId": fmt.Sprintf("%d", requestId)})
if err != nil {
	fmt.Println("Error:", err)
	return
}
requests, err := getgov.UnmarshalRequests(&body);
if err != nil {
	fmt.Println("Error:", err)
	return
}
// Create a table to display the response data
table := tablewriter.NewWriter(os.Stdout)
        table.SetHeader([]string{"ID", "Created", "GroupID", "TeamName", "Endpoint", "CallerID", "Status", "ProvisioningStep", "Message", "InitiatedBy", "Modified", "RetryCount", "QueuePriority"}) // Customize the table header as needed

        // Populate the table with data from the response
        for _, req := range requests {
            row := []string{
                fmt.Sprintf("%d", req.ID),
                req.Created,
                req.GroupID,
                req.TeamName,
                req.Endpoint,
                req.CallerID,
                req.Status,
                req.ProvisioningStep,
                req.Message,
                req.InitiatedBy,
                req.Modified,
                fmt.Sprintf("%v", req.RetryCount),
                fmt.Sprintf("%d", req.QueuePriority),
            
            }
            table.Append(row)
        }
        table.Render()
    },
}

func init() {

    getprovisioningstatusCmd.Flags().Int32VarP(&requestId, "requestId", "r", 0, "The id of the request to requeue")
  if err := requeueCmd.MarkFlagRequired("requestId"); err != nil {
		fmt.Println(err)
	}
    TeamGovCmd.AddCommand(getprovisioningstatusCmd)
}


