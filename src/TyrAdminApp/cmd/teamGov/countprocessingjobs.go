/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"github.com/spf13/cobra"
)

// countprocessingjobsCmd represents the countprocessingjobs command
var countprocessingjobsCmd = &cobra.Command{
	Use:   "countprocessingjobs",
	Short: "The commands counts the number of processing jobs in the Teams Governance API.",
	Long: `This command counts the number of processing jobs in the Teams Governance API. For example: 365Admin teamGov countprocessingjobs`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("help").Changed {
			cmd.Help()
		}
	},
}

func init() {
	TeamGovCmd.AddCommand(countprocessingjobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countprocessingjobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countprocessingjobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
