package graphhelper

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	// "github.com/microsoftgraph/msgraph-sdk-go/models"
	// "github.com/microsoftgraph/msgraph-sdk-go/users"
	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
	graphbeta "github.com/microsoftgraph/msgraph-beta-sdk-go"
)

type GraphHelper struct {
    clientSecretCredential *azidentity.ClientSecretCredential
    appClient              *msgraphsdk.GraphServiceClient
    betaClient             *graphbeta.GraphServiceClient
}

func NewGraphHelper() *GraphHelper {
    g := &GraphHelper{}
    return g
}

func (g *GraphHelper) InitializeGraphForAppAuth() error {
  	viper, err := viperConfig.InitViper("config.json")
    if err != nil {
        fmt.Printf("Error reading config file: %v\n", err)
    
    }

    clientId := viper.GetString("O365AzureAppClientId")
    tenantId := viper.GetString("O365TenantName")
    clientSecret := viper.GetString("O365AzureAppClientSecret")
  
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
    betaClient := graphbeta.NewGraphServiceClient(adapter)
    g.appClient = client
    g.betaClient = betaClient

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


// func listUsers(graphHelper *graphhelper.GraphHelper) {
//     users, err := *graphHelper.GetUsers()
//     if err != nil {
//         log.Panicf("Error getting users: %v", err)
//     }

//     // Output each user's details
//     for _, user := range users.GetValue() {
//         fmt.Printf("User: %s\n", *user.GetDisplayName())
//         fmt.Printf("  ID: %s\n", *user.GetId())

//         noEmail := "NO EMAIL"
//         email := user.GetMail()
//         if email == nil {
//             email = &noEmail
//         }
//         fmt.Printf("  Email: %s\n", *email)
//     }

//     // If GetOdataNextLink does not return nil,
//     // there are more users available on the server
//     nextLink := users.GetOdataNextLink()

//     fmt.Println()
//     fmt.Printf("More users available? %t\n", nextLink != nil)
//     fmt.Println()
// }