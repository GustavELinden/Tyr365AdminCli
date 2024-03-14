package teamGov

import (
	"fmt"
	"os"

	saveToFile "github.com/GustavELinden/TyrAdminCli/365Admin/SaveToFile"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)
var callerID string
// getfailedrequestsCmd represents the getfailedrequests command
var getfailedrequestsCmd = &cobra.Command{
	Use:   "getfailedrequests",
	Short: "Gets failed requests in the Teams Governance API by callerID",
	Long: `This command gets failed requests in the Teams Governance API by callerID. For example: 365Admin teamGov getfailedrequests.
		The response is a table with the following columns: ID, Created, GroupID, TeamName, Endpoint, CallerID, Status, ProvisioningStep, Message, InitiatedBy, Modified, RetryCount, QueuePriority.
		You specify the source by using the flag --callerID. For example: 365Admin teamGov getfailedrequests --callerID "Tyra".
	`,
	Run: func(cmd *cobra.Command, args []string) {
response, err :=	Get("GetFailedRequests", map[string]string{"callerID": callerID})
if err != nil {
	fmt.Println("Error:", err)
	return
}
requests, errs := UnmarshalRequests(&response)
if errs !=nil {
	fmt.Println("Error:", errs)
	return
}

if cmd.Flag("excel").Changed {
	var fileName string
	fmt.Println("Name your new excel file:")
	fmt.Scanln(&fileName)
	saveToFile.SaveToExcel(requests, fileName)
}
if cmd.Flag("print").Changed {
	renderRequests(requests)
}
if cmd.Flag("json").Changed {
	var fileName string
	fmt.Println("Enter a name for the JSON file (without extension):")
	fmt.Scanln(&fileName)

	err := saveToFile.SaveDataToJSONFile(requests, fileName+".json")
	if err != nil {
		fmt.Printf("Error saving data to JSON file: %s\n", err)
		return
	}
	fmt.Println("Data successfully saved to JSON file:", fileName+".json")
}
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
	getfailedrequestsCmd.Flags().StringVarP(&callerID, "callerID", "c", "", "The callerID to filter the failed requests")
    getfailedrequestsCmd.Flags().Bool("print", false, "Print the response as a table")
    getfailedrequestsCmd.Flags().Bool("excel", false, "Save the response to an Excel file")
    getfailedrequestsCmd.Flags().Bool("json", false, "Save the response to a JSON file")
	TeamGovCmd.AddCommand(getfailedrequestsCmd)

}
