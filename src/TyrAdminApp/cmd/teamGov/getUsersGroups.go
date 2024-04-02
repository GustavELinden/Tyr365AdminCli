/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userEmail string

// getUsersGroupsCmd represents the getUsersGroups command
var getUsersGroupsCmd = &cobra.Command{
	Use:   "getUsersGroups",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getUsersGroups called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	getUsersGroupsCmd.Flags().StringVarP(&userEmail, "userEmail", "u", "", "User Email")
	TeamGovCmd.AddCommand(getUsersGroupsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getUsersGroupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getUsersGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
