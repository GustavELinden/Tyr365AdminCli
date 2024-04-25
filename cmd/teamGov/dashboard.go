/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package teamGov

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	azurehelper "github.com/GustavELinden/Tyr365AdminCli/AzureHelper"
	teamGovHttp "github.com/GustavELinden/Tyr365AdminCli/TeamsGovernance"
	GraphHelper "github.com/GustavELinden/Tyr365AdminCli/graphHelper"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

type MatchRequest struct {
	Ids []string `json:"ids"`
}

var taskList []string
var deletedGroups []pterm.BulletListItem
var graphHelper *GraphHelper.GraphHelper
var metrics *azurehelper.MetricsResult
	var avgProvTime string
// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Renders a dashboard",
	Long: `Renders a dashboard with statistics about the goverance API and todos for me`,
	Run: func(cmd *cobra.Command, args []string) {

		taskList, _ = graphHelper.GetAllTasks()
		deletedGroups, _ = listDeletedGroups()
		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
    fmt.Print("Not logged into Azure. Metrics wont me used")
		}
		if cred != nil {
			metrics, _ = azurehelper.GetMetrics()
			if err != nil {
				fmt.Println(err)
			}
		}
		drawDashboard()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		go func() {
			for {
				select {
				case <-ticker.C:
					// Your update logic here
					// This is where you refresh or redraw your dashboard
					drawDashboard()
				case <-c:
					fmt.Println("Dashboard closed.")
					os.Exit(0)
				}
			}
		}()

		// Keep the main goroutine alive, or optionally do more work here
		select {}
	},
}

func init() {

	graphHelper = GraphHelper.NewGraphHelper()

	Initialize(graphHelper)
	TeamGovCmd.AddCommand(dashboardCmd)

}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Initialize(graphHelper *GraphHelper.GraphHelper) {
	err := graphHelper.InitializeGraphForAppAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for app auth: %v\n", err)
	}
}
func drawDashboard() {
	requests, _ := getprocessingjobs()
  avgProvTime, _ := getAverageProcessingTime()
 
	// Prepare table data and render table
	tableData, _ := buildRequestDataTable(requests)
	section2Table, _ := pterm.DefaultTable.WithHasHeader().WithData(tableData).Srender()
	clearScreen() // Assuming you have this function ready


	// Prepare bullet list
	bulletListItems := []pterm.BulletListItem{}
	if len(taskList) > 0 {
		for _, task := range taskList {
			bpoint := pterm.BulletListItem{Level: 0, Text: task}
			bulletListItems = append(bulletListItems, bpoint)
		}
	}
	section3List, _ := pterm.DefaultBulletList.WithItems(bulletListItems).Srender()

	section2ListMetrics, _ := buildAndRenderMetrics()

	section2List, _ := pterm.DefaultBulletList.WithItems(deletedGroups).Srender()

	panels := pterm.Panels{
		{
			{Data: pterm.DefaultSection.Sprint("TeamGov Status:")},
			{Data: pterm.DefaultSection.Sprint("TeamGov Metrics:")},
			{Data: pterm.DefaultSection.Sprint("TeamGov Average provisioning time.")},
		},
		// Second Row of Panels
		{
			{Data: section2Table},
			{Data: *section2ListMetrics},
			{Data: avgProvTime},
		},
		{

			{Data: pterm.DefaultSection.Sprint("Not-started Todos")},
			{Data: pterm.DefaultSection.Sprint("Deleted groups")},
		},
		// Second Row of Panels
		{
			{Data: section3List},
			{Data: section2List},
		},
	}

	// Adjust PanelPrinter settings if necessary to better fit your content
	panelPrinter := pterm.DefaultPanel.WithPanels(panels).WithPadding(20)
	panelPrinter.Padding = 20
	panelPrinter.SameColumnWidth = true

	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("365", pterm.NewStyle(pterm.FgLightCyan)),
		putils.LettersFromStringWithStyle("Admin", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)
	// render all things here
	_ = panelPrinter.Render()
}

func getprocessingjobs() (teamGovHttp.RequestSlice, error) {
	body, err := teamGovHttp.Get("GetProcessingJobs")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	requests, err := teamGovHttp.UnmarshalRequests(&body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return requests, nil
}

func listDeletedGroups() ([]pterm.BulletListItem, error) {
	groups, err := graphHelper.GetDeletedGroups()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var groupIds []string
	for _, group := range groups {
		groupIds = append(groupIds, *group.GetId())
	}

	requestBody := MatchRequest{
		Ids: groupIds,
	}

	body, err := teamGovHttp.PostWithBody("GetManagedTeams", nil, requestBody)
	if err != nil {
		fmt.Println("Failed to get managed teams:", err)
		return nil, err
	}
	var managedTeams []teamGovHttp.ManagedTeam
	err = json.Unmarshal(body, &managedTeams)

	if err != nil {
		fmt.Println("Failed to unmarshal managed teams:", err)
		return nil, err
	}

	bulletListItems := []pterm.BulletListItem{}

	// we add flag to print which team has which Origin and Retention
	for _, team := range managedTeams {
		if team.Origin == "GovPortal" && team.Retention == "Forever" {
			bpoint := pterm.BulletListItem{Level: 0, Text: team.TeamName + " is from " + team.Origin + " and needs to be discussed"}
			bulletListItems = append(bulletListItems, bpoint)

		}
		if team.Origin == "Tyra" && team.Retention == "Forever" {
			bpoint := pterm.BulletListItem{Level: 0, Text: team.TeamName + " is from " + team.Origin + " and needs to be discussed"}
			bulletListItems = append(bulletListItems, bpoint)
		} else {
			fmt.Println(team.TeamName + " is from " + team.Origin + " and does not need to be restored")
		}
	}
	return bulletListItems, nil
}

func getAverageProcessingTime()(string, error){
		body, err := teamGovHttp.Get("AverageProvisionTime")
		if err != nil {
			fmt.Println("Error:", err)
		return	 "", err 
		}

 return string(body), nil
}

func buildRequestDataTable(requests teamGovHttp.RequestSlice)(pterm.TableData, error){
tableData :=	pterm.TableData{
		{"RequestId", "TeamName", "EndPoint"},
	}

	for _, req := range requests {
		row := []string{
			fmt.Sprintf("%d", req.ID),
			req.TeamName,
			req.Endpoint,
		}
		tableData = append(tableData, row)
	}
	return tableData, nil
}

func buildAndRenderMetrics()(*string, error){
		bulletListMetrics := []pterm.BulletListItem{
		{Level: 0, Text: fmt.Sprintf("TeamGovernance avg responsetime: %f", metrics.AverageResponseTime)},
		{Level: 0, Text: fmt.Sprintf("TeamGovernance Http5xx requests (1h): %f", metrics.Http5xxCount)},
		{Level: 0, Text: fmt.Sprintf("TeamGovernance nr Requests (1h): %f", metrics.TotalRequests)},
	}

	section2ListMetrics, err := pterm.DefaultBulletList.WithItems(bulletListMetrics).Srender()
	if err != nil {
  return nil, err
	}
	return &section2ListMetrics, nil
}