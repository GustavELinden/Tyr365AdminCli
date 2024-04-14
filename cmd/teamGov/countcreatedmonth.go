/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"encoding/json"
	"fmt"

	"github.com/GustavELinden/Tyr365AdminCli/teamGovHttp"
	"github.com/spf13/cobra"
)

var Createdmonth int32

var countcreatedmonthCmd = &cobra.Command{
	Use:   "countcreatedmonth",
	Short: "Returns the number of requests created in the current month.",
	Long: `Returns the number of requests created in the current month. For example: 365Admin teamGov countcreatedmonth
	The endpoints counted against are the following: Create, ApplySPTemplate, ApplyTeamTemplate, Group`,
	Run: func(cmd *cobra.Command, args []string) {
		numbm, err := teamGovHttp.Get("CountEntriesMonth")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		errAtParsing := json.Unmarshal(numbm, &Createdmonth)
		if errAtParsing != nil {
			fmt.Println("Error:", errAtParsing)
			return
		}
		fmt.Println(Createdmonth)
	},
}

func init() {
	TeamGovCmd.AddCommand(countcreatedmonthCmd)
}
