package teamGovHttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	viperConfig "github.com/GustavELinden/Tyr365AdminCli/config"
	"github.com/olekukonko/tablewriter"
)

var TokenCache string

func makePOSTRequest(postUrl string, bodyValues []byte) (*http.Response, error) {
	// Encode the body values into a URL-encoded format
	body := bytes.NewBuffer(bodyValues)

	// Create the request
	req, err := http.NewRequest("POST", postUrl, body)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func AuthGovernanceApi() (string, error) {
	viper, err := viperConfig.InitViper("config.json")
	if err != nil {
		fmt.Printf("Error initializing viper: %v\n", err)
		return "", errors.New("error initializing viper")
	}
	authAdress := "https://login.microsoftonline.com/a2728528-eff8-409c-a379-7d900c45d9ba/oauth2/token"

	bodyValues := url.Values{}
	bodyValues.Set("grant_type", viper.GetString("grant_type"))
	bodyValues.Set("client_id", viper.GetString("client_id"))
	bodyValues.Set("client_secret", viper.GetString("client_secret"))
	bodyValues.Set("resource", viper.GetString("resource"))
	body := []byte(bodyValues.Encode())
	// Make the POST request
	resp, err := makePOSTRequest(authAdress, body)
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return "", errors.New("error making POST request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status code: %d\n", resp.StatusCode)
		return "", errors.New("unexpected response status code")
	}

	// Decode the response body
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		return "", errors.New("error decoding response body")
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		fmt.Println("Access token not found in response")
		return "", errors.New("access token not found in response")
	}
	// Print the access token

	return accessToken, nil
}
func AuthGraphApi() (string, error) {
	viper, err := viperConfig.InitViper("config.json")
	if err != nil {
		fmt.Printf("Error initializing viper: %v\n", err)
		return "", errors.New("error initializing viper")
	}

	authAdress := "https://login.microsoftonline.com/a2728528-eff8-409c-a379-7d900c45d9ba/oauth2/token"

	bodyValues := url.Values{}
	bodyValues.Set("grant_type", viper.GetString("grant_type"))
	bodyValues.Set("client_id", viper.GetString("M365managementAppClientId"))
	bodyValues.Set("client_secret", viper.GetString("M365ManagementAppClientSecret"))
	bodyValues.Set("resource", "https://graph.microsoft.com")
	body := []byte(bodyValues.Encode())
	// Make the POST request
	resp, err := makePOSTRequest(authAdress, body)
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return "", errors.New("error making POST request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status code: %d\n", resp.StatusCode)
		return "", errors.New("unexpected response status code")
	}

	// Decode the response body
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		return "", errors.New("error decoding response body")
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		fmt.Println("Access token not found in response")
		return "", errors.New("access token not found in response")
	}
	// Print the access token

	return accessToken, nil

}
func RetrieveAuthToken() (string, error) {
	return TokenCache, nil
}

func PrintToken() {
	fmt.Println(TokenCache)
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
	var teams []teamGovHttp.ManagedTeam
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
	req.Header.Set("Authorization", "Bearer "+token)
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
	req.Header.Set("Authorization", "Bearer "+token)
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

func UnmarshalRequests(body *[]byte) (RequestSlice, error) {
	var requests []teamGovHttp.Request
	err := json.Unmarshal(*body, &requests)
	if err == nil {
		return requests, nil
	}

	var request teamGovHttp.Request
	err = json.Unmarshal(*body, &request)
	if err == nil {
		return teamGovHttp.RequestSlice{request}, nil
	}

	return nil, fmt.Errorf("error unmarshalling to Request or []Request: %w", err)
}
func UnmarshalGroups(body *[]byte) ([]teamGovHttp.UnifiedGroup, error) {
	var groups []UnifiedGroup
	err := json.Unmarshal(*body, &groups)
	if err == nil {
		return groups, nil
	}

	var group UnifiedGroup
	err = json.Unmarshal(*body, &group)
	if err == nil {
		return []UnifiedGroup{group}, nil
	}

	return nil, fmt.Errorf("error unmarshalling to UnifiedGroup or []UnifiedGroup: %w", err)
}
func UnmarshalManagedTeams(body *[]byte) ([]teamGovHttp.ManagedTeam, error) {
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

func GetTaskETag(taskID string) (string, error) {
	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/planner/tasks/%s/details", taskID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	accessToken, err := AuthGraphApi()
	req.Header.Add("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return "", err
	}
	defer resp.Body.Close()

	// ETag is found in the "ETag" response

	etag := resp.Header.Get("ETag")
	if etag == "" {
		return "", fmt.Errorf("ETag header not found in response")
	}
	return etag, nil
}

type RequestSlice []teamGovHttp.Request
type Printer interface {
	PrintTable()
}
func(r *RequestSlice) PrintTable(){

	// Create a table to display the response data
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Created", "GroupID", "TeamName", "Endpoint", "CallerID", "Status", "ProvisioningStep", "Message", "InitiatedBy", "Modified", "RetryCount", "QueuePriority"}) // Customize the table header as needed
	for _, req := range *r {
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
	table.Render()
}