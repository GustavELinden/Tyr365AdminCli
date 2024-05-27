package teamGov

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	viperConfig "github.com/GustavELinden/Tyr365AdminCli/config"
	logging "github.com/GustavELinden/Tyr365AdminCli/logger"
	log "github.com/sirupsen/logrus"
)

type TokenCached struct {
	Token string
}

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
	logger := logging.GetLogger()

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
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"status":          "Error",
			"queryparameters": bodyValues,
		}).Error("error making POST request")

		return "", errors.New("error making POST request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"status":          "Error",
			"statusCode":      resp.StatusCode,
			"queryparameters": bodyValues,
		}).Error("error making POST request")
		return "", errors.New("unexpected response status code")
	}

	// Decode the response body
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"error":           err,
			"status":          "Error",
			"queryparameters": bodyValues,
		}).Error("error decoding token")
		return "", errors.New("error decoding response body")
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"status":          "Error",
			"queryparameters": bodyValues,
		}).Error("Token was empty")
		return "", errors.New("access token not found in response")
	}
	// Print the access token

	return accessToken, nil
}
func AuthGraphApi() (string, error) {
	logger := logging.GetLogger()
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
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"status":          "Error",
			"queryparameters": bodyValues,
		}).Error("Token was empty")
		return "", errors.New("error making POST request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"status":          resp.Status,
			"statusCode":      resp.StatusCode,
			"queryparameters": bodyValues,
		}).Error("Statuscode indicates error")
		return "", errors.New("unexpected response status code")
	}

	// Decode the response body
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		logger.WithFields(log.Fields{
			"url":    authAdress,
			"method": "POST",
			"status": resp.Status,

			"queryparameters": bodyValues,
		}).Error("could not decode token")
		return "", errors.New("error decoding response body")
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		logger.WithFields(log.Fields{
			"url":             authAdress,
			"method":          "POST",
			"status":          resp.Status,
			"statusCode":      resp.StatusCode,
			"queryparameters": bodyValues,
		}).Error("token was empty")
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
