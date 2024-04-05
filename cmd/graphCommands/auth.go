/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"
	"log"

	"github.com/GustavELinden/TyrAdminCli/cmd/teamGov"
	graphhelper "github.com/GustavELinden/TyrAdminCli/graphHelper"
	"github.com/spf13/cobra"
)

var graphHelper *graphhelper.GraphHelper

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")

		// displayAccessToken(graphHelper)
	token, err :=	teamGov.AuthGovernanceApi()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(token)
	},
}

func init() {
	graphHelper = graphhelper.NewGraphHelper()

	initializeGraph(graphHelper)
	GraphCmd.AddCommand(authCmd)

}
func initializeGraph(graphHelper *graphhelper.GraphHelper) {
	err := graphHelper.InitializeGraphForAppAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for app auth: %v\n", err)
	}
}
func displayAccessToken(graphHelper *graphhelper.GraphHelper) {
	token, err := graphHelper.GetAppToken()
	if err != nil {
		log.Panicf("Error getting user token: %v\n", err)
	}

	fmt.Printf("App-only token: %s", *token)
	fmt.Println()
}
