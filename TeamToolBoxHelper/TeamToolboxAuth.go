package teamToolboxHelper

import (
	"context"
	"fmt"
	"net/http"

	viperConfig "github.com/GustavELinden/Tyr365AdminCli/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// TokenProvider holds the configuration for OAuth2 token acquisition
type TokenProvider struct {
    ClientID     string
    ClientSecret string
    TenantID     string
    APIClientID  string
}

func CreateClient()(*APIClient, error){
viper, _ := viperConfig.InitViper("config.json")
tokenProvider := &TokenProvider{
        ClientID:     viper.GetString("365ManagementAppId"),
        ClientSecret: viper.GetString("365ManagementAppSecret"),
        TenantID:     viper.GetString("O365TenantName"),
        APIClientID:  viper.GetString("teamToolboxAppId"),
}
    apiClient := &APIClient{
        AuthProvider: tokenProvider,
        BaseURL:     viper.GetString("teamToolBoxAdress"),
    }

  return apiClient, nil
}

// GetToken returns an OAuth2 token using client credentials
func (tp *TokenProvider) GetToken() (*oauth2.Token, error) {
    config := clientcredentials.Config{
        ClientID:     tp.ClientID,
        ClientSecret: tp.ClientSecret,
        TokenURL:     fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tp.TenantID),
        Scopes:       []string{fmt.Sprintf("api://%s/.default", tp.APIClientID)},
    }

    ctx := context.Background()
    token, err := config.Token(ctx)
    if err != nil {
        return nil, fmt.Errorf("unable to get token: %w", err)
    }

    return token, nil
}

// GetAuthenticatedClient returns an HTTP client authenticated with the token
func (tp *TokenProvider) GetAuthenticatedClient() (*http.Client, error) {
    config := clientcredentials.Config{
        ClientID:     tp.ClientID,
        ClientSecret: tp.ClientSecret,
        TokenURL:     fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tp.TenantID),
        Scopes:       []string{fmt.Sprintf("api://%s/.default", tp.APIClientID)},
    }
// 
    ctx := context.Background()
    client := config.Client(ctx)
    return client, nil
}