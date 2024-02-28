/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var propertiesFlag []string
var amountFlag int32

var getusersCmd = &cobra.Command{
	Use:   "getusers",
	Short: "Retrieves users from Microsoft Graph",
	Long: `Retrieves a list of users from Microsoft Graph, with customizable properties via the command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		users, err := graphHelper.GetUsers(propertiesFlag, &amountFlag)
		if err != nil {
			fmt.Printf("Error getting users: %v\n", err)
			return
		}

		fmt.Println("Users retrieved successfully.")

		caser := cases.Title(language.English) // Create a Title caser for English
		for _, user := range users.GetValue() {
			userValue := reflect.ValueOf(user).Elem() // Assume user is a pointer; adjust if not
			for _, prop := range propertiesFlag {
				prop = caser.String(prop) // Use the Title caser here
				fieldValue := userValue.FieldByName(prop)
				if fieldValue.IsValid() {
					// Check if the field is a pointer and not nil
					if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
						// Dereference pointer
						fieldValue = fieldValue.Elem()
					}
					// Handle potential interface{}
					if fieldValue.Kind() == reflect.Interface && !fieldValue.IsNil() {
						fieldValue = fieldValue.Elem()
					}
					fmt.Printf("%s: %v\n", prop, fieldValue.Interface())
				} else {
					fmt.Printf("%s: property not found\n", prop)
				}
			}
			fmt.Println("---") // Separator between users
		}
	},
}

func init() {
	GraphCmd.AddCommand(getusersCmd)
	getusersCmd.Flags().StringSliceVarP(&propertiesFlag, "properties", "p", []string{"displayName", "id", "mail"}, "Properties to select (comma-separated)")
	getusersCmd.Flags().Int32VarP(&amountFlag, "amount", "a", 25, "Amount of users to retrieve")
}