/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var Createdweek int32

// countcreatedweekCmd represents the countcreatedweek command
var countcreatedweekCmd = &cobra.Command{
	Use:   "countcreatedweek",
	Short: "Count the number of created requests in the Teams Governance API for the current week.",
	Long: `This command counts the number of created requests in the Teams Governance API for the current week. For example: 365Admin teamGov countcreatedweek
	The endpoints counted against are the following: Create, ApplySPTemplate, ApplyTeamTemplate, Group`,
	Run: func(cmd *cobra.Command, args []string) {
		numbm, err := Get("CountEntriesWeek")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		errAtParsing := json.Unmarshal(numbm, &Createdweek)
		if errAtParsing != nil {
			fmt.Println("Error:", errAtParsing)
			return
		}
		fmt.Println(Createdweek)
	},
}

func init() {
	TeamGovCmd.AddCommand(countcreatedweekCmd)

}
