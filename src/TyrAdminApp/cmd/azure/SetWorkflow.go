package azure

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)
var status string
// disableWorkflowCmd represents the disableWorkflow command
var setWorkflowCmd = &cobra.Command{
	Use:   "setWorkflow",
	Short: "Starts or Stops the Teams Governance API trigger.",
	Long: `This command starts or stops the Teams Governance API trigger. For example: 365Admin azure setWorkflow --status "Enable"`,
	Run: func(cmd *cobra.Command, args []string) {
        status, _ := cmd.Flags().GetString("status")
        if status != "Enabled" && status != "Disabled" {
            fmt.Println("Error: --status must be 'Enabled' or 'Disabled'")
            return
        }

        // Construct the Azure CLI command with the status flag
        azureCLICommand := fmt.Sprintf("az logicapp config appsettings set --name tyrensteamsgovprocessrequests --resource-group teamsprovisioning --settings Workflows.ProcessRequestsWorkflow.FlowState=%s", status)

        // Execute the Azure CLI command
        if output, err := executeAzureCLICommand(azureCLICommand); err != nil {
            fmt.Printf("Failed to execute Azure CLI command: %v\n", err)
        } else {
            fmt.Println("Command Output:", output)
        }
    },
}


func init() {
    setWorkflowCmd.Flags().StringVarP(&status, "status", "s", "", "The status of the Teams Governance API Logicapp trigger (e.g., 'Enabled' or 'Disabled')")
	AzureCmd.AddCommand(setWorkflowCmd)
}

func executeAzureCLICommand(command string) (string, error) {
    parts := strings.Fields(command)
    cmd := exec.Command(parts[0], parts[1:]...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("error executing command: %v, output: %s", err, string(output))
    }
    return string(output), nil
}


