package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

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

func GetTokenForGovernanceApi() string {
	authAdress := "https://login.microsoftonline.com/a2728528-eff8-409c-a379-7d900c45d9ba/oauth2/token"

	// Define the body properties
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
		return ""
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status code: %d\n", resp.StatusCode)
		return ""
	}

	// Decode the response body
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		return ""
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		fmt.Println("Access token not found in response")
		return ""
	}

	// Print the access token
	fmt.Println("Access Token:", accessToken)
  return accessToken
}
