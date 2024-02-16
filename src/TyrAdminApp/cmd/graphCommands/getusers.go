/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"

	"github.com/spf13/cobra"
)
var propertiesFlag []string
var amountFlag int32
// getusersCmd represents the getusers command
var getusersCmd = &cobra.Command{
	Use:   "getusers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
users, err := graphHelper.GetUsers(propertiesFlag, &amountFlag)
        if err != nil {
            fmt.Printf("Error getting users: %v\n", err)
            return
        }
        // Handle `users` response
        fmt.Println("Users retrieved successfully.")
	// Output each user's details
	for _, user := range users.GetValue() {
		fmt.Printf("User: %s\n", *user.GetDisplayName())
		fmt.Printf("  ID: %s\n", *user.GetId())

		noEmail := "NO EMAIL"
		email := user.GetMail()
		if email == nil {
			email = &noEmail
		}
		fmt.Printf("  Email: %s\n", *email)
	}
	},
}

func init() {
	GraphCmd.AddCommand(getusersCmd)
	getusersCmd.Flags().StringSliceVarP(&propertiesFlag, "properties", "p", []string{"displayName", "id", "mail"}, "Properties to select (comma-separated)")
	getusersCmd.Flags().Int32VarP(&amountFlag, "amount", "a", 25, "Amount of users to retrieve")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getusersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getusersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
