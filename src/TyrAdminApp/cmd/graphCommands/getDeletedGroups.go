/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package graphCommands

import (
	"encoding/json"
	"fmt"

	"github.com/GustavELinden/TyrAdminCli/365Admin/cmd/teamGov"
	"github.com/spf13/cobra"
)

//Make struct for the bmodels.Group
type MatchRequest struct {
    Ids []string `json:"ids"`
}

var getDeletedGroupsCmd = &cobra.Command{
	Use:   "getDeletedGroups",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
  groups, err := graphHelper.GetDeletedGroups()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    

    var groupIds []string
    for _, group := range groups {
        groupIds = append(groupIds, *group.GetId())
    }


    requestBody := MatchRequest{
        Ids: groupIds,
    }
    
    body, err := teamGov.PostWithBody("GetManagedTeams", nil, requestBody)
    if err != nil {
        fmt.Println("Failed to get managed teams:", err)
        return
    }
    var managedTeams []teamGov.ManagedTeam
    err = json.Unmarshal(body, &managedTeams)
	
    if err != nil {
        fmt.Println("Failed to unmarshal managed teams:", err)
        return
    }


for _, team := range managedTeams {
    if team.Origin == "GovPortal" || team.Origin == "Tyra" && team.Retention == "Forever"{
        fmt.Println(team.TeamName + " is from " + team.Origin + " and needs to be restored")
    } else {
        fmt.Println(team.TeamName + " is from " + team.Origin + " and does not need to be restored")
    }
}


},
}


func init() {
	GraphCmd.AddCommand(getDeletedGroupsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDeletedGroupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDeletedGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

