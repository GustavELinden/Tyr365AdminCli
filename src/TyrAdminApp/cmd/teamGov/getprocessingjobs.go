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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getprocessingjobs called")
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
