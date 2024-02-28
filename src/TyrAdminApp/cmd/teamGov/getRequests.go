package teamGov

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

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
type Parameters struct {
    GroupID        string `json:"groupId"`
    TemplateId     int    `json:"templateId"`
    Description    string `json:"description"`
    CallerId       string `json:"callerId"`
    InitiatedBy    string `json:"initiatedBy"`
    FlowParameters struct {
        TemplateID        string `json:"templateID"`
        ProjectNumber     string `json:"ProjectNumber"`
        TyrAProcessType   string `json:"TyrAProcessType"`
    } `json:"flowParameters"`
    ClientTaskId   int `json:"clientTaskId"`
    // Add other fields as needed
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
type ManagedTeam struct {
	Id        int    `json:"Id"`
	GroupId   string `json:"groupId"`
	TeamName  string `json:"teamName"`
	Status    string `json:"status"`
	Origin    string `json:"origin"`
	Retention string `json:"retention"`
}

func Get(endpoint string, queryParams ...map[string]string) ([]byte, error) {

    viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        return nil, fmt.Errorf("failed to initialize viper: %w", err)
    }

    token, err := AuthGovernanceApi()

  

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
    req.Header.Set("Authorization", "Bearer "+ token)

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

func GetQuery(targetEndpoint string, queryParams map[string]string) ([]byte, error) {
    // Initialize Viper to load configuration, assuming viperConfig is correctly set up in your project.
    viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        return nil, fmt.Errorf("failed to initialize viper: %w", err)
    }

    // Retrieve the authentication token.
    token, err := AuthGovernanceApi()
    if err != nil || token == "" {
        return nil, fmt.Errorf("failed to get authentication token: %v", err)
    }

    // Construct the API URL.
    apiURL := viper.GetString("resource") + "/api/teams/" + targetEndpoint
    query := url.Values{}

    // Iterate over the map of query parameters and add them to the query string.
    for key, value := range queryParams {
        query.Set(key, value) // Use Set or Add depending on your needs.
    }

    // Append the query string to the API URL.
    if len(query) > 0 {
        apiURL += "?" + query.Encode()
    }
    fmt.Println(apiURL);
    // Create the HTTP GET request.
    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+token)

    // Execute the HTTP request.
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    // Check the response status code.
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
    }

    // Read and return the response body.
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    return body, nil
}

func QueryManaged(groupId, teamName, status, origin, retention, fields string) ([]ManagedTeam, error) {
	// Initialize Viper to load configuration
	viper, err := viperConfig.InitViper("config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize viper: %w", err)
	}

	// Retrieve the authentication token
	token, err := AuthGovernanceApi()
	if err != nil || token == "" {
		return nil, fmt.Errorf("failed to get authentication token: %v", err)
	}

	// Construct the API URL and query parameters
	apiURL := viper.GetString("resource") + "/api/teams/query"
	query := url.Values{}
	if groupId != "" {
		query.Set("groupId", groupId)
	}
	if teamName != "" {
		query.Set("teamName", teamName)
	}
	if status != "" {
		query.Set("status", status)
	}
	if origin != "" {
		query.Set("origin", origin)
	}
	if retention != "" {
		query.Set("retention", retention)
	}
	if fields != "" {
		query.Set("fields", fields)
	}

	// Append query string to the API URL
	if len(query) > 0 {
		apiURL += "?" + query.Encode()
	}

	// Create the HTTP GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Read and unmarshal the response body
	var teams []ManagedTeam
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	err = json.Unmarshal(body, &teams)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return teams, nil
}

func Post(endpoint string, queryParams map[string]string) ([]byte, error) {
    viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        return nil, fmt.Errorf("failed to initialize viper: %w", err)
    }

    token, err := AuthGovernanceApi()
    if err != nil || token == "" {
        return nil, errors.New("failed to get authentication token")
    }

    apiURL := viper.GetString("resource") + "/api/teams/" + endpoint
    if len(queryParams) > 0 {
        query := url.Values{}
        for key, value := range queryParams {
            query.Add(key, value)
        }
        apiURL += "?" + query.Encode()
    }

    req, err := http.NewRequest("POST", apiURL, nil) // No body is sent
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+ token)
    // No need to set Content-Type for a request without a body

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
    }

    response, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    return response, nil
}
func PostWithBody(endpoint string, queryParams map[string]string, body interface{}) ([]byte, error) {
    viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        return nil, fmt.Errorf("failed to initialize viper: %w", err)
    }

     token, err := AuthGovernanceApi()
    if err != nil || token == "" {
        return nil, errors.New("failed to get authentication token")
    }

    apiURL := viper.GetString("resource") + "/api/teams/" + endpoint
    if len(queryParams) > 0 {
        query := url.Values{}
        for key, value := range queryParams {
            query.Add(key, value)
        }
        apiURL += "?" + query.Encode()
    }

    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, fmt.Errorf("error marshalling request body: %w", err)
    }

    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+ token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
    }

    response, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    return response, nil
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
func UnmarshalManagedTeams(body *[]byte) ([]ManagedTeam, error) {
    var managedTeam []ManagedTeam
    err := json.Unmarshal(*body, &managedTeam)
    if err == nil {
        return managedTeam, nil
    }

    var request ManagedTeam
    err = json.Unmarshal(*body, &request)
    if err == nil {
        return []ManagedTeam{request}, nil
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