package GraphHelper

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/GustavELinden/Tyr365AdminCli/internal/config"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

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
	cfg := config.Get()

	clientId := cfg.GetString("M365managementAppClientId")
	tenantId := cfg.GetString("O365TenantName")
	clientSecret := cfg.GetString("M365ManagementAppClientSecret")

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
