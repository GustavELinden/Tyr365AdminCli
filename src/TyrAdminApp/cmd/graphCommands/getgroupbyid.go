/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getgroupbyidCmd represents the getgroupbyid command
var getgroupbyidCmd = &cobra.Command{
	Use:   "getgroupbyid",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
 _, err := graphHelper.GetGroupById("12c53ea7-9b08-4b4e-a1d7-06c8ced53d8c")
	 if err != nil {		
		fmt.Println("Error:", err)
		return
	}

	 
	},
}

func init() {
	GroupsCmd.AddCommand(getgroupbyidCmd)
 
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getgroupbyidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getgroupbyidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}