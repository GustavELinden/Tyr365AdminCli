package teamGov

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Assuming these variables are declared at the package level to store flag values
var (
    endpoint       string
    created        string
    createdEnd     string
    callerId       string
    initiatedByUser string

    top            int // Assuming there's a sensible default or 0 indicates "use default"
)

// newCmd represents the command for the new endpoint
var queryCmd = &cobra.Command{
    Use:   "query",
    Short: "Description of what the new command does",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Processing flags and constructing query parameters map
        queryParams := make(map[string]string)
        if endpoint != "" {
            queryParams["endpoint"] = endpoint
        }
        if created != "" {
            queryParams["created"] = created
        }
        if createdEnd != "" {
            queryParams["createdEnd"] = createdEnd
        }
        if callerId != "" {
            queryParams["callerId"] = callerId
        }
        if initiatedByUser != "" {
            queryParams["initiatedByUser"] = initiatedByUser
        }
        if status != "" {
            queryParams["status"] = status
        }
        if top > 0 { // Assuming a non-zero value should be included
            queryParams["top"] = fmt.Sprintf("%d", top)
        }

        body, err := GetQuery("CliQuery", queryParams)
       requests, err := UnmarshalRequests(&body);
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
}

func init() {
    // Register flags for the newCmd
    queryCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "Comma-separated endpoints")
    queryCmd.Flags().StringVarP(&created, "created", "c", "", "Start date (YYYY/MM/DD)")
    queryCmd.Flags().StringVarP(&createdEnd, "createdEnd", "C", "", "End date (YYYY/MM/DD)")
    queryCmd.Flags().StringVarP(&callerId, "callerId", "i", "", "Comma-separated caller IDs")
   queryCmd.Flags().StringVarP(&initiatedByUser, "initiatedBy", "u", "", "User who initiated")
    queryCmd.Flags().StringVarP(&status, "status", "s", "", "Comma-separated statuses")
    queryCmd.Flags().IntVarP(&top, "top", "t", 0, "Limit the number of results")

    TeamGovCmd.AddCommand(queryCmd) // Assuming TeamGovCmd is your root or sub-root command
}
