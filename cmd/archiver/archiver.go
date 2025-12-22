/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package archivercmd

import "github.com/spf13/cobra"

// archiverCmd represents the archiver command
var ArchiverCmd = &cobra.Command{
	Use:   "archiver",
	Short: "Interact with the M365 Archiver orchestrator",
	Long:  "Collection of commands that wrap the M365 Archiver orchestrator API operations.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
