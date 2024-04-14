package tblprinter

import (
	"fmt"
	"os"

	"github.com/GustavELinden/Tyr365AdminCli/teamGovHttp"
	"github.com/olekukonko/tablewriter"
)

func RenderTable(requests []teamGovHttp.Request) {
	// Reflect the slice to work with its elements
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
}
