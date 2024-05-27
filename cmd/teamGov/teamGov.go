package teamGov

import (
	"github.com/spf13/cobra"
)

var (
	requestId int32  // Assuming requestId is declared here as global as well
	status    string // Global variable declaration for status
)

// teamGovCmd represents the teamGov command
var TeamGovCmd = &cobra.Command{
	Use:   "teamGov",
	Short: "teamGov is a palett that contains commands to manage the Teams Governance API.",
	Long: `teamGov is a palett that contains commands to manage the Teams Governance API. 
	For example: 365Admin teamGov countprocessingjobs`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
