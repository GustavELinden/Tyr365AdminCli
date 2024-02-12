package getgov

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/GustavELinden/TyrAdminCli/365Admin/authentication"
	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
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
    ExpirationDateTime string   `json:"expirationDateTime"` // Assuming expirationDateTime is a date/time field
    ExchangeProperties interface{} `json:"exchangeProperties"`
}

// unmarshalResponse tries to unmarshal JSON into either a single Request object or a slice of Request objects
// func unmarshalResponse(body []byte) ([]Request, error) {
//     var requests []Request
//     if err := json.Unmarshal(body, &requests); err == nil {
//         return requests, nil
//     }

//     var singleRequest Request
//     if err := json.Unmarshal(body, &singleRequest); err == nil {
//         return []Request{singleRequest}, nil
//     }

//     return nil, errors.New("response did not match expected JSON structures")
// }
func Get(endpoint string, queryParams ...map[string]string) ([]byte, error) {
    viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        return nil, fmt.Errorf("failed to initialize viper: %w", err)
    }

    token, err := authentication.GetAuthToken()
    if token == "" {
        return nil, errors.New("failed to get authentication token")
    }
    if err != nil {
        return nil, errors.New("failed to get authentication token")
    }


    apiURL := viper.GetString("resource") + "/api/teams/" + endpoint
    if len(queryParams) > 0 {
        query := url.Values{}
        for key, value := range queryParams[0] {
            query.Add(key, value)
        }
        apiURL += "?" + query.Encode()
    }

    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    return body, nil
}
func UnmarshalRequests(body *[]byte) ([]Request, error) {
    var requests []Request
    err := json.Unmarshal(*body, &requests)
    if err == nil {
        return requests, nil
    }

    var request Request
    err = json.Unmarshal(*body, &request)
    if err == nil {
        return []Request{request}, nil
    }

    return nil, fmt.Errorf("error unmarshalling to Request or []Request: %w", err)
}

func UnmarshalInteger(body *[]byte) (int, error) {
    var value int
    err := json.Unmarshal(*body, &value)
    if err != nil {
        return 0, fmt.Errorf("error unmarshalling to integer: %w", err)
    }
    return value, nil
}

func PrintJSONResponseString(body *[]byte) error {
    var responseString string
    fmt.Println(*body)
    err := json.Unmarshal(*body, &responseString)
    if err != nil {
        return fmt.Errorf("error unmarshalling response body to string: %w", err)
    }
    fmt.Println(&responseString)
    return nil
}

// Add more unmarshal functions as needed for other response types

// func GetRequests(endpoint string, queryParams ...map[string]string) ([]Request, error) {
//     viper, err := viperConfig.InitViper("config.json")
//     if err != nil {
//         return nil, fmt.Errorf("failed to initialize viper: %w", err)
//     }

//     token := authentication.GetTokenForGovernanceApi()
//     if token == "" {
//         return nil, errors.New("failed to get authentication token")
//     }

//     apiURL := viper.GetString("resource") + "/api/teams/" + endpoint

//     if len(queryParams) > 0 {
//         query := url.Values{}
//         for key, value := range queryParams[0] {
//             query.Add(key, value)
//         }
//         apiURL += "?" + query.Encode()
//     }

//     req, err := http.NewRequest("GET", apiURL, nil)
//     if err != nil {
//         return nil, fmt.Errorf("error creating request: %w", err)
//     }

//     req.Header.Set("Authorization", "Bearer "+token)

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         return nil, fmt.Errorf("error sending request: %w", err)
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != http.StatusOK {
//         return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
//     }

//     body, err := io.ReadAll(resp.Body)
//     if err != nil {
//         return nil, fmt.Errorf("error reading response body: %w", err)
//     }

//     // Use the unmarshalResponse helper function to handle both potential response types
//     requests, err := unmarshalResponse(body)
//     if err != nil {
//         return nil, err
//     }

//     return requests, nil
// }