/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)
var CreatedYear int32
// countcreatedCmd represents the countcreated command
var countcreatedCmd = &cobra.Command{
	Use:   "countcreatedyear",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	numbm, err :=	Get("CountEntriesYear")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
   errAtParsing := json.Unmarshal(numbm, &CreatedYear)
	 if errAtParsing != nil {
		fmt.Println("Error:", errAtParsing)
		return
	 }
   fmt.Println(CreatedYear)
	},
}

func init() {
	TeamGovCmd.AddCommand(countcreatedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countcreatedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countcreatedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
