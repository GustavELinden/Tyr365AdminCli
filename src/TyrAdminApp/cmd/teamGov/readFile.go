/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"encoding/json"
	"fmt"

	saveToFile "github.com/GustavELinden/TyrAdminCli/365Admin/SaveToFile"
	"github.com/spf13/cobra"
)
	var fileName string
// readFileCmd represents the readFile command
var readFileCmd = &cobra.Command{
	Use:   "readFile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var readGroups interface {}
	
		// fmt.Println("Enter the name of the file you want to read from: ")
		// fmt.Scan(&fileName)
    saveToFile.ReadDataFromJSONFile(fileName, &readGroups)
		// prettyJSON, _ := json.MarshalIndent(readGroups, "", "    ")

	    if cmd.Flag("output").Changed {
             outData,_ := json.Marshal(readGroups)
            fmt.Println(string(outData))
	};
},
}

func init() {
	readFileCmd.Flags().StringVarP(&fileName, "file", "f", "", "The name of the file you want to read from")
	TeamGovCmd.AddCommand(readFileCmd)
}
