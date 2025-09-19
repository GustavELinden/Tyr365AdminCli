package archiver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	viperConfig "github.com/GustavELinden/Tyr365AdminCli/config"
)

// apiToken represents a cached API token with expiration
type apiToken struct {
	Token     string
	ExpiresAt time.Time
}

var (
	tokenCache = make(map[string]apiToken)
	tokenMutex sync.RWMutex
)

// Thread-safe token cache access functions
func getTokenFromCache(resource string) (apiToken, bool) {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	token, found := tokenCache[resource]
	return token, found
}

func setTokenInCache(resource string, token apiToken) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	tokenCache[resource] = token
}

// makePOSTRequest creates and executes a POST request
func makePOSTRequest(postUrl string, bodyValues []byte) (*http.Response, error) {
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

// AuthArchiverApi gets an authentication token for the Archiver API
func AuthArchiverApi() (string, error) {
	viper, err := viperConfig.InitViper("config.json")
	if err != nil {
		return "", fmt.Errorf("error initializing viper: %w", err)
	}
	
	resource := viper.GetString("archiverAdress")
	if resource == "" {
		return "", errors.New("archiverAdress not found in configuration")
	}

	// Check cache first
	if token, found := getTokenFromCache(resource); found && time.Now().Before(token.ExpiresAt) {
		return token.Token, nil
	}

	authAddress := "https://login.microsoftonline.com/a2728528-eff8-409c-a379-7d900c45d9ba/oauth2/token"

	bodyValues := url.Values{}
	bodyValues.Set("grant_type", "client_credentials")
	bodyValues.Set("client_id", viper.GetString("archiverApp"))
	bodyValues.Set("client_secret", viper.GetString("archiverSecret"))
	bodyValues.Set("resource", resource)
	body := []byte(bodyValues.Encode())

	// Make the POST request
	resp, err := makePOSTRequest(authAddress, body)
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	// Decode the response body
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("error decoding response body: %w", err)
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", errors.New("access token not found in response")
	}

	// Handle expires_in as both string and number
	var expiresIn int
	if expiresInFloat, ok := tokenResponse["expires_in"].(float64); ok {
		expiresIn = int(expiresInFloat)
	} else if expiresInStr, ok := tokenResponse["expires_in"].(string); ok {
		var err error
		expiresIn, err = strconv.Atoi(expiresInStr)
		if err != nil {
			return "", fmt.Errorf("could not convert expires_in to integer: %w", err)
		}
	} else {
		// Default to 1 hour if expires_in is missing
		expiresIn = 3600
	}

	// Cache the token with 5-minute buffer
	setTokenInCache(resource, apiToken{
		Token:     accessToken,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(expiresIn-300)), // 5-min buffer
	})

	return accessToken, nil
}

// Data models based on OpenAPI schemas

// ArchiveJob represents an archive job
type ArchiveJob struct {
	ID                           int             `json:"id"`
	IDExportDataJobs             int             `json:"idExportDataJobs"`
	FilePath                     string          `json:"filePath"`
	GroupID                      string          `json:"groupId"`
	Status                       *string         `json:"status"`
	SharePointURL                *string         `json:"sharePointUrl"`
	Created                      *time.Time      `json:"created"`
	Alias                        string          `json:"alias"`
	ArchiveSubJobs               []ArchiveSubJob `json:"archiveSubJobs"`
	IDExportDataJobsNavigation   interface{}     `json:"idExportDataJobsNavigation"`
}

// ArchiveSubJob represents a sub-job within an archive job
type ArchiveSubJob struct {
	ID                         int         `json:"id"`
	IDArchiveJobs              int         `json:"idArchiveJobs"`
	Type                       string      `json:"type"`
	Status                     *string     `json:"status"`
	GroupID                    *string     `json:"groupId"`
	Created                    *time.Time  `json:"created"`
	IDArchiveJobsNavigation    *ArchiveJob `json:"idArchiveJobsNavigation"`
}

// ExportDataJob represents an export data job
type ExportDataJob struct {
	ID          int          `json:"id"`
	GroupID     string       `json:"groupId"`
	Alias       string       `json:"alias"`
	Status      string       `json:"status"`
	FilePath    string       `json:"filePath"`
	SiteURL     string       `json:"siteUrl"`
	RequestID   int          `json:"requestId"`
	ArchiveJobs []ArchiveJob `json:"archiveJobs"`
	Request     *Request     `json:"request"`
}

// Request represents a request
type Request struct {
	ID                int           `json:"id"`
	Created           time.Time     `json:"created"`
	GroupID           *string       `json:"groupId"`
	TeamName          *string       `json:"teamName"`
	Endpoint          *string       `json:"endpoint"`
	CallerID          *string       `json:"callerId"`
	Parameters        *string       `json:"parameters"`
	Status            *string       `json:"status"`
	ProvisioningStep  *string       `json:"provisioningStep"`
	Message           *string       `json:"message"`
	InitiatedBy       *string       `json:"initiatedBy"`
	Modified          *time.Time    `json:"modified"`
	RowVersion        []byte        `json:"rowVersion"`
	ClientTaskID      *int          `json:"clientTaskId"`
	LTPMessageSent    *bool         `json:"ltpmessageSent"`
	Hidden            *bool         `json:"hidden"`
	GroupInformation  *string       `json:"groupInformation"`
	RetryCount        *int          `json:"retryCount"`
	QueuePriority     *int          `json:"queuePriority"`
	ExportDataJobs    []ExportDataJob `json:"exportDataJobs"`
	RequestSteps      []RequestStep   `json:"requestSteps"`
}

// RequestStep represents a step in a request
type RequestStep struct {
	ID        int       `json:"id"`
	RequestID int       `json:"requestId"`
	Step      *string   `json:"step"`
	Status    *string   `json:"status"`
	Message   *string   `json:"message"`
	Modified  *time.Time `json:"modified"`
	Request   *Request  `json:"request"`
}

// ArchiverClient provides methods to interact with the M365 Archiver API
type ArchiverClient struct {
	baseURL string
}

// NewArchiverClient creates a new archiver client
func NewArchiverClient() (*ArchiverClient, error) {
	viper, err := viperConfig.InitViper("config.json")
	if err != nil {
		return nil, fmt.Errorf("error initializing viper: %w", err)
	}
	
	baseURL := viper.GetString("archiverAdress")
	if baseURL == "" {
		return nil, errors.New("archiverAdress not found in configuration")
	}

	return &ArchiverClient{
		baseURL: baseURL,
	}, nil
}

// makeAuthenticatedRequest makes an HTTP request with authentication
func (c *ArchiverClient) makeAuthenticatedRequest(method, endpoint string, body io.Reader, queryParams map[string]string) ([]byte, error) {
	// Get authentication token
	token, err := AuthArchiverApi()
	if err != nil {
		return nil, fmt.Errorf("failed to get authentication token: %w", err)
	}

	// Construct URL
	apiURL := c.baseURL + endpoint
	if len(queryParams) > 0 {
		query := url.Values{}
		for key, value := range queryParams {
			query.Set(key, value)
		}
		apiURL += "?" + query.Encode()
	}

	// Create request
	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return responseBody, nil
}

// GET endpoints

// CreateArchiveJobs calls GET /api/Archiver/CreateArchiveJobs
func (c *ArchiverClient) CreateArchiveJobs() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/CreateArchiveJobs", nil, nil)
}

// CreateArchiveSubJobs calls GET /api/Archiver/CreateArchiveSubJobs
func (c *ArchiverClient) CreateArchiveSubJobs() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/CreateArchiveSubJobs", nil, nil)
}

// GetJobs calls GET /api/Archiver/GetJobs
func (c *ArchiverClient) GetJobs() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/GetJobs", nil, nil)
}

// ProcessArchiveSubJobs calls GET /api/Archiver/ProcessArchiveSubJobs
func (c *ArchiverClient) ProcessArchiveSubJobs() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/ProcessArchiveSubJobs", nil, nil)
}

// ProcessStatusOfSubJobs calls GET /api/Archiver/ProcessStatusOfSubJobs
func (c *ArchiverClient) ProcessStatusOfSubJobs() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/ProcessStatusOfSubJobs", nil, nil)
}

// GetArchiveJobsByStatus calls GET /api/Archiver/GetArchiveJobsByStatus
func (c *ArchiverClient) GetArchiveJobsByStatus(status string) ([]byte, error) {
	params := map[string]string{"status": status}
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/GetArchiveJobsByStatus", nil, params)
}

// GetArchiveSubJobsByStatus calls GET /api/Archiver/GetArchiveSubJobsByStatus
func (c *ArchiverClient) GetArchiveSubJobsByStatus(status string) ([]byte, error) {
	params := map[string]string{"status": status}
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/GetArchiveSubJobsByStatus", nil, params)
}

// GetArchiveJobByID calls GET /api/Archiver/GetArchiveJobById
func (c *ArchiverClient) GetArchiveJobByID(id int) ([]byte, error) {
	params := map[string]string{"id": strconv.Itoa(id)}
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/GetArchiveJobById", nil, params)
}

// ReprocessSubJobs calls GET /api/Archiver/ReprocessSubJobs
func (c *ArchiverClient) ReprocessSubJobs() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/ReprocessSubJobs", nil, nil)
}

// GetArchiveJobDTO calls GET /api/Archiver/GetArchiveJobDTO
func (c *ArchiverClient) GetArchiveJobDTO(groupID string) ([]byte, error) {
	params := map[string]string{"groupId": groupID}
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/GetArchiveJobDTO", nil, params)
}

// FinishExport calls GET /api/Archiver/FinishExport
func (c *ArchiverClient) FinishExport() ([]byte, error) {
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/FinishExport", nil, nil)
}

// Test calls GET /api/Archiver/Test
func (c *ArchiverClient) Test(id int, groupID string) ([]byte, error) {
	params := map[string]string{
		"id":      strconv.Itoa(id),
		"groupId": groupID,
	}
	return c.makeAuthenticatedRequest("GET", "/api/Archiver/Test", nil, params)
}

// POST endpoints

// CreateArchiveJobbet calls POST /api/Archiver/CreateArchiveJobbet
func (c *ArchiverClient) CreateArchiveJobbet(archiveJob ArchiveJob) ([]byte, error) {
	jsonBody, err := json.Marshal(archiveJob)
	if err != nil {
		return nil, fmt.Errorf("error marshalling archive job: %w", err)
	}
	
	return c.makeAuthenticatedRequest("POST", "/api/Archiver/CreateArchiveJobbet", bytes.NewBuffer(jsonBody), nil)
}

// UpdateArchiveJob calls POST /api/Archiver/UpdateArchiveJob
func (c *ArchiverClient) UpdateArchiveJob(id int, status string) ([]byte, error) {
	params := map[string]string{
		"id":     strconv.Itoa(id),
		"status": status,
	}
	return c.makeAuthenticatedRequest("POST", "/api/Archiver/UpdateArchiveJob", nil, params)
}

// UpdateArchiveSubJob calls POST /api/Archiver/UpdateArchiveSubJob
func (c *ArchiverClient) UpdateArchiveSubJob(id int, status string) ([]byte, error) {
	params := map[string]string{
		"id":     strconv.Itoa(id),
		"status": status,
	}
	return c.makeAuthenticatedRequest("POST", "/api/Archiver/UpdateArchiveSubJob", nil, params)
}

// DELETE endpoints

// DeleteArchiveJob calls DELETE /api/Archiver/DeleteArchiveJob
func (c *ArchiverClient) DeleteArchiveJob(id int) ([]byte, error) {
	params := map[string]string{"id": strconv.Itoa(id)}
	return c.makeAuthenticatedRequest("DELETE", "/api/Archiver/DeleteArchiveJob", nil, params)
}

// DeleteArchiveSubJob calls DELETE /api/Archiver/DeleteArchiveSubJob
func (c *ArchiverClient) DeleteArchiveSubJob(id int) ([]byte, error) {
	params := map[string]string{"id": strconv.Itoa(id)}
	return c.makeAuthenticatedRequest("DELETE", "/api/Archiver/DeleteArchiveSubJob", nil, params)
}

// Helper functions for unmarshalling responses

// UnmarshalArchiveJob unmarshals a single archive job from response body
func UnmarshalArchiveJob(body []byte) (*ArchiveJob, error) {
	var job ArchiveJob
	err := json.Unmarshal(body, &job)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling archive job: %w", err)
	}
	return &job, nil
}

// UnmarshalArchiveJobs unmarshals multiple archive jobs from response body
func UnmarshalArchiveJobs(body []byte) ([]ArchiveJob, error) {
	var jobs []ArchiveJob
	err := json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling archive jobs: %w", err)
	}
	return jobs, nil
}

// UnmarshalArchiveSubJob unmarshals a single archive sub job from response body
func UnmarshalArchiveSubJob(body []byte) (*ArchiveSubJob, error) {
	var subJob ArchiveSubJob
	err := json.Unmarshal(body, &subJob)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling archive sub job: %w", err)
	}
	return &subJob, nil
}

// UnmarshalArchiveSubJobs unmarshals multiple archive sub jobs from response body
func UnmarshalArchiveSubJobs(body []byte) ([]ArchiveSubJob, error) {
	var subJobs []ArchiveSubJob
	err := json.Unmarshal(body, &subJobs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling archive sub jobs: %w", err)
	}
	return subJobs, nil
}

// UnmarshalExportDataJob unmarshals a single export data job from response body
func UnmarshalExportDataJob(body []byte) (*ExportDataJob, error) {
	var job ExportDataJob
	err := json.Unmarshal(body, &job)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling export data job: %w", err)
	}
	return &job, nil
}

// UnmarshalExportDataJobs unmarshals multiple export data jobs from response body
func UnmarshalExportDataJobs(body []byte) ([]ExportDataJob, error) {
	var jobs []ExportDataJob
	err := json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling export data jobs: %w", err)
	}
	return jobs, nil
}

// UnmarshalString unmarshals a simple string response
func UnmarshalString(body []byte) (string, error) {
	var result string
	err := json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling string response: %w", err)
	}
	return result, nil
}

// UnmarshalInteger unmarshals a simple integer response
func UnmarshalInteger(body []byte) (int, error) {
	var result int
	err := json.Unmarshal(body, &result)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling integer response: %w", err)
	}
	return result, nil
}