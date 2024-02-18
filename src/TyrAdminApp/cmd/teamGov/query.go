package teamGov

import (
	"fmt"

	saveToFile "github.com/GustavELinden/TyrAdminCli/365Admin/SaveToFile"
	tblprinter "github.com/GustavELinden/TyrAdminCli/365Admin/tblPrinter"
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
    Short: "Querys the governance API for requests",
    Long: `Querys the governance API for requests based on the provided parameters.
    The results can be printed as a table or saved to an Excel file.
    If no parameters are provided, the command will return the 50 latest Create requests
    
    Available parameters:
    --endpoint: Comma-separated endpoints (e.g. "endpoint1,endpoint2") . Endpoints are "Create", "ApplySPTemplate", "ApplyTeamTemplate", "Group", "ArchiveTeam". If no endpoint is provided, default endpoint is Create
    --created: Start date (YYYY/MM/DD) (e.g. "2021/01/01"). If no date is provided, default date is 60 days ago.
    --createdEnd: End date (YYYY/MM/DD) (e.g. "2021/01/01"). If no date is provided, default date is today.
    --callerId: Comma-separated caller IDs (e.g. "callerId1,callerId2"). Default callerId is "Tyra".
    --initiatedBy: User who initiated the request (e.g. "user1@tyrens.se"). If no user is provided, default user is "sposervice@tyrens.onmicrosoft.com".
    --status: Comma-separated statuses (e.g. "status1,status2"). Default status is "Succeeded". Available statuses are "Succeeded", "Error", "Queued", "Processing".
    --top: Limit the number of results. Default is 50. Max is 1000.
        .`,
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
       if cmd.Flag("excel").Changed {
        var fileName string
        fmt.Println("Name your new excel file:")
        fmt.Scanln(&fileName)
        saveToFile.SaveToExcel(requests, fileName)
    } 
    if cmd.Flag("print").Changed {
        tblprinter.RenderTable(requests)
    }
if err != nil {
	fmt.Println("Error:", err)
	return
}


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
    queryCmd.Flags().Bool("help", false, "Print help for the command")  
    queryCmd.Flags().Bool("excel", false, "Save the response to an Excel file")
    queryCmd.Flags().Bool("print", false, "Print the response as a table")
    TeamGovCmd.AddCommand(queryCmd) // Assuming TeamGovCmd is your root or sub-root command
}
