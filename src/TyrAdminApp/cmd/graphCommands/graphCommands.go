/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"fmt"
	"log"

	graphhelper "github.com/GustavELinden/TyrAdminCli/365Admin/graphHelper"
	"github.com/spf13/cobra"
)

// testGraphCmd represents the testGraph command
var TestGraphCmd = &cobra.Command{
	Use:   "testGraph",
	Short: "Creates a graph client and gets an app-only token.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	graphHelper := graphhelper.NewGraphHelper()

    initializeGraph(graphHelper)
		displayAccessToken(graphHelper)
	
	},
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
func init() {
 
}
