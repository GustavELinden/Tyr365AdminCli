package graphhelper

import (
	"context"

	// users  "github.com/microsoftgraph/msgraph-sdk-go/users"
	models "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	//other-imports
)

func (g *GraphHelper) GetGroupById(groupId string) (models.Groupable, error) {
   group, err := g.appClient.Groups().ByGroupId(groupId).Get(context.Background(), nil)
	 if err != nil {
		 return nil, err
	 }			
	
	 return group, nil
}

func (g *GraphHelper) GetUsers(selectProperties []string, amount *int32) (models.UserCollectionResponseable, error) {
    var topValue int32
    if amount == nil {
        topValue = 25 // Default value if amount is not provided
    } else {
        topValue = *amount
    }

    query := users.UsersRequestBuilderGetQueryParameters{
        Select:  selectProperties,
        Top:     &topValue,
        Orderby: []string{"displayName"},
    }

    return g.appClient.Users().
        Get(context.Background(), &users.UsersRequestBuilderGetRequestConfiguration{
            QueryParameters: &query,
        })
}

// func printGroup(group models.Groupable) {
//     // Marshaling the group to JSON for a more readable output
//     groupJson, err := json.MarshalIndent(group, "", "  ")
//     if err != nil {
//         fmt.Println("Error marshaling group:", err)
//         return
//     }
//     fmt.Println(string(groupJson))
// }