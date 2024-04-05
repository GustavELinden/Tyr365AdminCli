package teamGov

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var CreatedYear int32

var countcreatedCmd = &cobra.Command{
	Use:   "countcreatedyear",
	Short: "Counts the number of Requests created this calendar year.",
	Long: `This command will return the number of requests created this calendar year. For example: 365Admin teamGov countcreatedyear
	The endpoints counted against are the following: Create, ApplySPTemplate, ApplyTeamTemplate, Group`,
	Run: func(cmd *cobra.Command, args []string) {
		numbm, err := Get("CountEntriesYear")
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
}
