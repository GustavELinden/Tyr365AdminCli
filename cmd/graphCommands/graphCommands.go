/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"github.com/spf13/cobra"
)

// testGraphCmd represents the testGraph command
var GraphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Creates a graph client and gets an app-only token.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

}
