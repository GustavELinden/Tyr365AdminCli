/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"
	"os"

	getgov "github.com/GustavELinden/TyrAdminCli/365Admin/httpFuncs"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// getprocessingjobsCmd represents the getprocessingjobs command
var getprocessingjobsCmd = &cobra.Command{
	Use:   "getprocessingjobs",
	Short: "Get the processing jobs in the Teams Governance API.",
	Long: `This command gets the processing jobs in the Teams Governance API. For example: 365Admin teamGov getprocessingjobs. 
    The response is a table with the following columns: ID, Created, GroupID, TeamName, Endpoint, CallerID, Status, ProvisioningStep, Message, InitiatedBy, Modified, RetryCount, QueuePriority.`,
	Run: func(cmd *cobra.Command, args []string) {
				if cmd.Flag("help").Changed {
			cmd.Help()
		}
		body, err := getgov.Get("GetProcessingJobs")
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

        // Render the table
        table.Render()
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        // Parse flags before the command runs
        err := cmd.Flags().Parse(args)
        if err != nil {
            fmt.Println("Error parsing flags:", err)
        }
    
	},
}

func init() {

  TeamGovCmd.AddCommand(getprocessingjobsCmd)

}
