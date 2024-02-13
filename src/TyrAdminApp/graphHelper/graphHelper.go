package graphhelper

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	// "github.com/microsoftgraph/msgraph-sdk-go/models"
	// "github.com/microsoftgraph/msgraph-sdk-go/users"

	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
)

type GraphHelper struct {
    clientSecretCredential *azidentity.ClientSecretCredential
    appClient              *msgraphsdk.GraphServiceClient
}

func NewGraphHelper() *GraphHelper {
    g := &GraphHelper{}
    return g
}

func (g *GraphHelper) InitializeGraphForAppAuth() error {
  	viper, err := viperConfig.InitViper("config.json")

    clientId := os.Getenv(viper.GetString("O365AzureAppClientId"))
    tenantId := os.Getenv(viper.GetString("O365TenantName"))
    clientSecret := os.Getenv(viper.GetString("O365AzureAppClientSecret"))
    credential, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
    if err != nil {
        return err
    }

    g.clientSecretCredential = credential

    // Create an auth provider using the credential
    authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(g.clientSecretCredential, []string{
        "https://graph.microsoft.com/.default",
    })
    if err != nil {
        return err
    }

    // Create a request adapter using the auth provider
    adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
    if err != nil {
        return err
    }

    // Create a Graph client using request adapter
    client := msgraphsdk.NewGraphServiceClient(adapter)
    g.appClient = client

    return nil
}

func (g *GraphHelper) GetAppToken() (*string, error) {
    token, err := g.clientSecretCredential.GetToken(context.Background(), policy.TokenRequestOptions{
        Scopes: []string{
            "https://graph.microsoft.com/.default",
        },
    })
    if err != nil {
        return nil, err
    }

    return &token.Token, nil
}