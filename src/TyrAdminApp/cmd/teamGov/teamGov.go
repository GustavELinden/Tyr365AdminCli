package teamGov

import (
	"github.com/spf13/cobra"
)

// teamGovCmd represents the teamGov command
var TeamGovCmd = &cobra.Command{
	Use:   "teamGov",
	Short: "teamGov is a palett that contains commands to manage the Teams Governance API.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()	
	},
}

func init() {


}
