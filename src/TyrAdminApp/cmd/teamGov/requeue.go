/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)
var (
	requestId int32 
)
// requeueCmd represents the requeue command
var requeueCmd = &cobra.Command{
	Use:   "requeue",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("requeue called")
        //Print the requestId from the command line
        
        fmt.Println(requestId)
   
	},
}

func init() {
	requeueCmd.Flags().Int32VarP(&requestId, "requestId", "r", 0, "The id of the request to requeue")
  if err := requeueCmd.MarkFlagRequired("requestId"); err != nil {
		fmt.Println(err)
	}

  TeamGovCmd.AddCommand(requeueCmd)

}
