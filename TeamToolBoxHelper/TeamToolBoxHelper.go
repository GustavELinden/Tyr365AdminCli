package teamToolboxHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
)

func (client *APIClient) GetTestPolicy() (string, error) {
    httpClient, err := client.AuthProvider.GetAuthenticatedClient()
    if err != nil {
        return "", err
    }

    resp, err := httpClient.Get(fmt.Sprintf("%s/Tools/TestPolicy", client.BaseURL))
    if err != nil {
        return "", fmt.Errorf("unable to call API: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("unable to read response: %w", err)
    }

    return string(body), nil
}

func (client *APIClient) GetToolById(id string) (*TblTool, error) {
      httpClient, err := client.AuthProvider.GetAuthenticatedClient()
    if err != nil {
        return nil, err
    }
    var adress string = fmt.Sprintf("%s/Admin/" + id, client.BaseURL)
    resp, err := httpClient.Get(adress)
    if err != nil {
        return nil, fmt.Errorf("unable to call public API: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
   
    if err != nil {
        return nil, fmt.Errorf("unable to read public response: %w", err)
    }
    var tool TblTool
    err = json.Unmarshal(body, &tool)
    if err != nil {
        fmt.Println(err)
    }
    return &tool, nil
}

func (client *APIClient) GetRulesAndLogic() (*RulesandLogics, error) {
      httpClient, err := client.AuthProvider.GetAuthenticatedClient()
    if err != nil {
        return nil, err
    }
    var adress string = fmt.Sprintf("%s/Admin/RulesandLogic", client.BaseURL)
    resp, err := httpClient.Get(adress)
    if err != nil {
        return nil, fmt.Errorf("unable to call public API: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
   
    if err != nil {
        return nil, fmt.Errorf("unable to read public response: %w", err)
    }
    var rulesandLogic RulesandLogics
    err = json.Unmarshal(body, &rulesandLogic)
    if err != nil {
        fmt.Println(err)
    }
    return &rulesandLogic, nil
}

func (client *APIClient) PostWithJSONBody(endpoint string, jsonBody []byte) ([]byte, error) {
	ttclient, err := client.AuthProvider.GetAuthenticatedClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticated client: %w", err)
	}

	address := fmt.Sprintf("%s/Admin/%s", client.BaseURL, endpoint)

	resp, err := ttclient.Post(address, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return response, nil
}




func (r *TblTool) PrintTable() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "ToolName", "CurrentTemplateId", "TopicName"}) 

		row := []string{
			fmt.Sprintf("%d", r.Id),
			r.ToolName,
           	fmt.Sprintf("%d", r.CurrentTempateId),
            r.TopicName,
		}
		table.Append(row)
	
	table.Render()
}

func (r *TblTools) PrintTable() {

	// Create a table to display the response data
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "ToolName", "CurrentTemplateId", "TopicName"}) // Customize the table header as needed
for _, req := range *r {
		row := []string{
			fmt.Sprintf("%d", req.Id),
			req.ToolName,
           	fmt.Sprintf("%d", req.CurrentTempateId),
            req.TopicName,
		}
		table.Append(row)
    }
	
	table.Render()
}

func (r *RulesandLogics) PrintTable() {

	// Create a table to display the response data
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "RuleName", "ToolId", "RuleId", "Value", "Logic"}) // Customize the table header as needed
for _, req := range *r {
		row := []string{
			fmt.Sprintf("%d", req.Id),
			req.RuleName,
           	fmt.Sprintf("%d", req.ToolId),
            fmt.Sprintf("%d", req.RuleId),
            req.Value,
            req.Logic,
		}
		table.Append(row)
    }
	
	table.Render()
}

func MarshalToJSON(body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}
	return jsonBody, nil
}