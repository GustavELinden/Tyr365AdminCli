/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	saveToFile "github.com/GustavELinden/TyrAdminCli/365Admin/SaveToFile"
	"github.com/spf13/cobra"
)

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "deleteFile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		//if flag --file is used, delete the file
		if cmd.Flag("file").Changed {
			err := saveToFile.DeleteFile(fileName)
			if err != nil {
				fmt.Printf("Error deleting file: %s\n", err)
				return
			}
			fmt.Println("File deleted successfully")
		}
	},
}

func init() {
	TeamGovCmd.AddCommand(deleteFileCmd)
  deleteFileCmd.Flags().StringVarP(&fileName, "file", "f", "", "The name of the file you want to delete")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
