package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
)
type TokenCache struct {
    Token     string
    ExpiresAt time.Time
}
var tokenCache *TokenCache

func (t *TokenCache) IsValid() bool {
    return time.Now().Before(t.ExpiresAt)
}


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

func getTokenForGovernanceApi() string {
	viper, err := viperConfig.InitViper("config.json")

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
	fmt.Println("Access Token aquired" )
  return accessToken
}

func GetAuthToken() (string, error) {
    if tokenCache != nil && tokenCache.IsValid() {
			fmt.Println("Token is valid")
        return tokenCache.Token, nil
    }

    // Your existing auth logic here
    // Assume newToken and expiresIn are obtained after authentication
    newToken := getTokenForGovernanceApi()
    expiresIn := 10 * time.Minute // Example duration

    tokenCache = &TokenCache{
        Token:     newToken,
        ExpiresAt: time.Now().Add(expiresIn),
    }

    return newToken, nil
}
