/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"time"

	"github.com/spf13/cobra"
)
type Group struct {
	ID                        *string    `json:"id,omitempty"`
	DeletedDateTime           *time.Time `json:"deletedDateTime,omitempty"`
	CreatedDateTime           *time.Time `json:"createdDateTime,omitempty"`
	CreatedByAppId            *string    `json:"createdByAppId,omitempty"`
	OrganizationId            *string    `json:"organizationId,omitempty"`
	Description               *string    `json:"description,omitempty"`
	DisplayName               *string    `json:"displayName,omitempty"`
	GroupTypes                []*string  `json:"groupTypes,omitempty"`
	InfoCatalogs              []*string  `json:"infoCatalogs,omitempty"`
	Mail                      *string    `json:"mail,omitempty"`
	MailEnabled               *bool      `json:"mailEnabled,omitempty"`
	MailNickname              *string    `json:"mailNickname,omitempty"`
	ProxyAddresses            []*string  `json:"proxyAddresses,omitempty"`
	RenewedDateTime           *time.Time `json:"renewedDateTime,omitempty"`
	Visibility                *string    `json:"visibility,omitempty"`
	
}
// groupsCmd represents the groups command
var GroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	
	GraphCmd.AddCommand(GroupsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
