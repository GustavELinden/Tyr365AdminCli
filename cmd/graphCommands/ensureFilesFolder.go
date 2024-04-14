/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"

	"github.com/spf13/cobra"
)
var groupId string
// ensureFilesFolderCmd represents the ensureFilesFolder command
var ensureFilesFolderCmd = &cobra.Command{
	Use:   "ensureFilesFolder",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	group, _ :=	graphHelper.GetGroupById(groupId)
  teamId := group.GetId()
	channels, _ := graphHelper.GetAllChannels(*teamId)
	allChannels := channels.GetValue()
for _, channel := range allChannels {
	chanId := channel.GetId() // Declare with :=
	chanTitle := channel.GetDisplayName() // Declare with :=

	fmt.Println("Ensure files folders called for: " + *chanTitle)

	_, err := graphHelper.EnsureFilesFolder(*teamId, *chanId)
	if err != nil {
		fmt.Println("Error calling:", *chanTitle)
		continue // Use continue to skip the current iteration on error
	}
	fmt.Println("Successfully made call")
}

   },
	}





func init() {
		ensureFilesFolderCmd.Flags().StringVarP(&groupId, "groupId", "r", "", "The id of the group")
	if err := ensureFilesFolderCmd.MarkFlagRequired("groupId"); err != nil {
		fmt.Println(err)
	}
	GraphCmd.AddCommand(ensureFilesFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ensureFilesFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ensureFilesFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
