package getgov

import (
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"

	"github.com/GustavELinden/TyrAdminCli/365Admin/authentication"
	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
)


func Get(endpoint string, queryParams ...map[string]string) {
viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        panic(err)
    }
		fmt.Println(viper.ConfigFileUsed())
    token := authentication.GetTokenForGovernanceApi()
    if token == "" {
        fmt.Println("Failed to get authentication token")
        return
    }

    // Construct the API URL
    apiURL := viper.GetString("resource") + "/api/teams/" + endpoint

    // Add query parameters to the URL, if provided
    if len(queryParams) > 0 {
        query := url.Values{}
        for key, value := range queryParams[0] {
            query.Add(key, value)
        }
        apiURL += "?" + query.Encode()
    }

    // Create a new HTTP GET request
    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Add bearer token to the request header
    req.Header.Set("Authorization", "Bearer "+token)

    // Create an HTTP client
    client := &http.Client{}

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    // Check the response status code
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error:", resp.Status)
        // Optionally, you might want to handle different status codes differently
        // For example, you might return an error or take some other action
        // based on the specific status code.
        return
    }

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    // Print the response body
    fmt.Println("Response:", string(body))
}

