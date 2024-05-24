/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"fmt"

	teamGovHttp "github.com/GustavELinden/Tyr365AdminCli/TeamsGovernance"
	"github.com/spf13/cobra"
)
var alias string
// AddFieldToSiteCmd represents the AddFieldToSite command
var AddFieldToSiteCmd = &cobra.Command{
	Use:   "AddFieldToSite",
	Short: "Calls the AddFieldToSite endpoint with alias parameter",
	Long: `The cmd calls an endpoint in the Teams Governance API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AddFieldToSite called")
		if cmd.Flag("alias").Changed {
			  queryParams := make(map[string]string)
        queryParams["alias"] = alias
				    _, err := teamGovHttp.Get("AddFieldToSite", queryParams)
                
                if err != nil {
                    fmt.Printf("Failed to add field %s: %v\n", err)
                    return
                }

                
                    fmt.Printf("Successfully added field to site %s\n", alias)
		}
	},
}

func init() {
	AddFieldToSiteCmd.Flags().StringVarP(&alias, "alias", "a", "", "alias of sp site")
	TeamGovCmd.AddCommand(AddFieldToSiteCmd)

}
