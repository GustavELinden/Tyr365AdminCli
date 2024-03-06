/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package azure

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		err := AuthenticateWithAzureCLI()
    if err != nil {
        fmt.Println("Authentication failed:", err)
        return
    }
	},
}

func init() {
	AzureCmd.AddCommand(loginCmd)

}
// AuthenticateWithAzureCLI initiates Azure CLI authentication.
func AuthenticateWithAzureCLI() error {
    fmt.Println("Initiating Azure login...")

    // Execute "az login" command
    cmd := exec.Command("az", "login")

    // Optionally, capture the output or error for logging purposes
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error during Azure CLI login: %v, output: %s", err, string(output))
    }

    fmt.Println("Login successful.")
    test()
    return nil
}

func test()error {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        panic("Failed to obtain a credential: " + err.Error())
    }

    // Example: Use the credential to list all subscriptions for the authenticated account
    subscriptionClient, err := armsubscriptions.NewClient(cred, nil)
    if err != nil {
        panic("Failed to create subscription client: " + err.Error())
    }

    pager := subscriptionClient.NewListPager(nil)
    for pager.More() {
        resp, err := pager.NextPage(context.Background())
        if err != nil {
            panic("Failed to list subscriptions: " + err.Error())
        }
        for _, subscription := range resp.Value {
            println(*subscription.SubscriptionID)
        }
    }
    return nil
}   
