package tblprinter

import (
	"fmt"
	"os"

	"github.com/GustavELinden/TyrAdminCli/365Admin/cmd/teamGov"
	"github.com/olekukonko/tablewriter"
)

type Request struct {
    ID                 int         `json:"Id"`
    Created            string      `json:"Created"`
    GroupID            string      `json:"GroupId"`
    TeamName           string      `json:"TeamName"`
    Endpoint           string      `json:"Endpoint"`
    CallerID           string      `json:"CallerId"`
    Parameters         string      `json:"Parameters"`
    Status             string      `json:"Status"`
    ProvisioningStep   string      `json:"ProvisioningStep"`
    Message            string      `json:"Message"`
    InitiatedBy        string      `json:"InitiatedBy"`
    Modified           string      `json:"Modified"`
    ClientTaskID       int         `json:"ClientTaskId"`
    LtpmessageSent     bool        `json:"LtpmessageSent"`
    Hidden             bool        `json:"Hidden"`
    RetryCount         int         `json:"RetryCount"`
    QueuePriority      int         `json:"QueuePriority"`
    GroupDetails       GroupDetails `json:"GroupDetails"`
}

type GroupDetails struct {
    GroupID           string    `json:"groupId"`
    DisplayName       string    `json:"displayName"`
    Alias             string    `json:"alias"`
    Description       string    `json:"description"`
    CreatedDate       string    `json:"createdDate"`
    SharePointURL     string    `json:"sharePointUrl"`
    Visibility        string    `json:"visibility"`
    Team              string    `json:"team"`
    Yammer            string    `json:"yammer"`
    Label             string    `json:"label"`
    ExpirationDateTime string   `json:"expirationDateTime"` 
    ExchangeProperties interface{} `json:"exchangeProperties"`
}

func RenderTable(requests []teamGov.Request) {
	// Reflect the slice to work with its elements
	table := tablewriter.NewWriter(os.Stdout)
        table.SetHeader([]string{"ID", "Created", "GroupID", "TeamName", "Endpoint", "CallerID", "Status", "ProvisioningStep", "Message", "InitiatedBy", "Modified", "RetryCount", "QueuePriority"}) // Customize the table header as needed

        // Populate the table with data from the response
        for _, req := range requests {
            row := []string{
                fmt.Sprintf("%d", req.ID),
                req.Created,
                req.GroupID,
                req.TeamName,
                req.Endpoint,
                req.CallerID,
                req.Status,
                req.ProvisioningStep,
                req.Message,
                req.InitiatedBy,
                req.Modified,
                fmt.Sprintf("%v", req.RetryCount),
                fmt.Sprintf("%d", req.QueuePriority),
            
            }
            table.Append(row)
        }

        // Render the table
        table.Render()
    }
