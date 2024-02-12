/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	getgov "github.com/GustavELinden/TyrAdminCli/365Admin/httpFuncs"
	"github.com/spf13/cobra"
)

// getprocessingrequestsCmd represents the getprocessingrequests command
var getprocessingrequestsCmd = &cobra.Command{
	Use:   "getprocessingrequests",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := getgov.Get("GetProcessingRequests")
if err != nil {
	fmt.Println("Error:", err)
	return
}
requests, err := getgov.UnmarshalInteger(&body);
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println(requests)
	},
}
// Create a table to display the response data
// table := tablewriter.NewWriter(os.Stdout)
//         table.SetHeader([]string{"ID", "Created", "GroupID", "TeamName", "Endpoint", "CallerID", "Status", "ProvisioningStep", "Message", "InitiatedBy", "Modified", "RetryCount", "QueuePriority"}) // Customize the table header as needed

//         // Populate the table with data from the response
//         for _, req := range requests {
//             row := []string{
//                 fmt.Sprintf("%d", req.ID),
//                 req.Created,
//                 req.GroupID,
//                 req.TeamName,
//                 req.Endpoint,
//                 req.CallerID,
//                 req.Status,
//                 req.ProvisioningStep,
//                 req.Message,
//                 req.InitiatedBy,
//                 req.Modified,
//                 fmt.Sprintf("%v", req.RetryCount),
//                 fmt.Sprintf("%d", req.QueuePriority),
            
//             }
//             table.Append(row)
//         }
//         table.Render()
    
// 	},
// }

func init() {
	TeamGovCmd.AddCommand(getprocessingrequestsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getprocessingrequestsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getprocessingrequestsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
