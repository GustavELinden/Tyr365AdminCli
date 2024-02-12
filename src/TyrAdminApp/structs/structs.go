package reusableStructs

import (
	"time"
)

// GroupDetails represents the nested "GroupDetails" object in the JSON
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
    ExpirationDateTime time.Time `json:"expirationDateTime"` // Assuming expirationDateTime is a date/time field
    ExchangeProperties interface{} `json:"exchangeProperties"`
}

// APIResponse represents the structure of the JSON object
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
    RowVersion         string      `json:"RowVersion"`
    ClientTaskID       int         `json:"ClientTaskId"`
    LtpmessageSent     bool        `json:"LtpmessageSent"`
    Hidden             bool        `json:"Hidden"`
    GroupInformation   interface{} `json:"GroupInformation"`
    RetryCount         interface{} `json:"RetryCount"`
    QueuePriority      int         `json:"QueuePriority"`
    ExportDataJobs     []interface{} `json:"ExportDataJobs"`
    GroupDetails       GroupDetails `json:"GroupDetails"`
}
