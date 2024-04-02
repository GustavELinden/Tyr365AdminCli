package graphhelper

import (
	"context"

	// users  "github.com/microsoftgraph/msgraph-sdk-go/users"
	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
	bmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	models "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	//other-imports
)
type User struct {
    ID                                   string   `json:"id,omitempty"`
    DeletedDateTime                      string `json:"deletedDateTime,omitempty"`
    AccountEnabled                       bool     `json:"accountEnabled,omitempty"`
    AgeGroup                             string   `json:"ageGroup,omitempty"`
    BusinessPhones                       []string `json:"businessPhones,omitempty"`
    City                                 string   `json:"city,omitempty"`
    CreatedDateTime                      string   `json:"createdDateTime,omitempty"`
    CreationType                         *string  `json:"creationType,omitempty"`
    CompanyName                          string   `json:"companyName,omitempty"`
    ConsentProvidedForMinor              *string  `json:"consentProvidedForMinor,omitempty"`
    Country                              string   `json:"country,omitempty"`
    Department                           string   `json:"department,omitempty"`
    DisplayName                          string   `json:"displayName,omitempty"`
    EmployeeId                           *string  `json:"employeeId,omitempty"`
    EmployeeHireDate                     *string  `json:"employeeHireDate,omitempty"`
    EmployeeLeaveDateTime                *string  `json:"employeeLeaveDateTime,omitempty"`
    EmployeeType                         *string  `json:"employeeType,omitempty"`
    FaxNumber                            *string  `json:"faxNumber,omitempty"`
    GivenName                            string   `json:"givenName,omitempty"`
    ImAddresses                          []string `json:"imAddresses,omitempty"`
    InfoCatalogs                         []string `json:"infoCatalogs,omitempty"`
    IsLicenseReconciliationNeeded        bool     `json:"isLicenseReconciliationNeeded,omitempty"`
    IsManagementRestricted               *bool    `json:"isManagementRestricted,omitempty"`
    IsResourceAccount                    *bool    `json:"isResourceAccount,omitempty"`
    JobTitle                             string   `json:"jobTitle,omitempty"`
    LegalAgeGroupClassification          string   `json:"legalAgeGroupClassification,omitempty"`
    Mail                                 string   `json:"mail,omitempty"`
    MailNickname                         string   `json:"mailNickname,omitempty"`
    MobilePhone                          string   `json:"mobilePhone,omitempty"`
    OnPremisesDistinguishedName          *string  `json:"onPremisesDistinguishedName,omitempty"`
    OfficeLocation                       string   `json:"officeLocation,omitempty"`
    OnPremisesDomainName                 *string  `json:"onPremisesDomainName,omitempty"`
    OnPremisesImmutableId                *string  `json:"onPremisesImmutableId,omitempty"`
    OnPremisesLastSyncDateTime           *string  `json:"onPremisesLastSyncDateTime,omitempty"`
    OnPremisesObjectIdentifier           *string  `json:"onPremisesObjectIdentifier,omitempty"`
    OnPremisesSecurityIdentifier         *string  `json:"onPremisesSecurityIdentifier,omitempty"`
    OnPremisesSamAccountName             *string  `json:"onPremisesSamAccountName,omitempty"`
    OnPremisesSyncEnabled                *bool    `json:"onPremisesSyncEnabled,omitempty"`
    OnPremisesUserPrincipalName          *string  `json:"onPremisesUserPrincipalName,omitempty"`
    OtherMails                           []string `json:"otherMails,omitempty"`
    PasswordPolicies                     *string  `json:"passwordPolicies,omitempty"`
    PostalCode                           string   `json:"postalCode,omitempty"`
    PreferredDataLocation                *string  `json:"preferredDataLocation,omitempty"`
    PreferredLanguage                    string   `json:"preferredLanguage,omitempty"`
    ProxyAddresses                       []string `json:"proxyAddresses,omitempty"`
    RefreshTokensValidFromDateTime       string   `json:"refreshTokensValidFromDateTime,omitempty"`
    SecurityIdentifier                   string   `json:"securityIdentifier,omitempty"`
    ShowInAddressList                    *bool    `json:"showInAddressList,omitempty"`
    SignInSessionsValidFromDateTime      string   `json:"signInSessionsValidFromDateTime,omitempty"`
    State                                string   `json:"state,omitempty"`
    StreetAddress                        string   `json:"streetAddress,omitempty"`
    Surname                              string   `json:"surname,omitempty"`
    UsageLocation                        string   `json:"usageLocation,omitempty"`
    UserPrincipalName                    string   `json:"userPrincipalName,omitempty"`
    ExternalUserConvertedOn              *string  `json:"externalUserConvertedOn,omitempty"`
    ExternalUserState                    *string  `json:"externalUserState,omitempty"`
    ExternalUserStateChangeDateTime      *string  `json:"externalUserStateChangeDateTime,omitempty"`
    UserType                             string   `json:"userType,omitempty"`
    EmployeeOrgData                      *string  `json:"employeeOrgData,omitempty"`
  
}
func (g *GraphHelper) GetGroupById(groupId string) (models.Groupable, error) {
   group, err := g.appClient.Groups().ByGroupId(groupId).Get(context.Background(), nil)
	 if err != nil {
		 return nil, err
	 }			
	
	 return group, nil
}

func (g *GraphHelper) GetDeletedGroups()([]bmodels.Groupable, error){
    graphGroups, err := g.betaClient.Directory().DeletedItems().GraphGroup().Get(context.Background(), nil)
    if err != nil {
        return nil, err
    }
    groups := graphGroups.GetValue()
    return groups, nil
}

func (g *GraphHelper) GetUsers(selectProperties []string, amount *int32,filter string ) (models.UserCollectionResponseable, error) {
    var topValue int32
    if amount == nil {
        topValue = 25 // Default value if amount is not provided
    } else {
        topValue = *amount
    }
 
    query := users.UsersRequestBuilderGetQueryParameters{
        Select:  selectProperties,
        Top:     &topValue,
        // Orderby: []string{"displayName"},
        Filter:  &filter,
      
    }

    return g.appClient.Users().
        Get(context.Background(), &users.UsersRequestBuilderGetRequestConfiguration{
            QueryParameters: &query,
        })
}

func (g *GraphHelper) CreateTask(taskTitle string) (models.PlannerTaskable, error) {
	// Initialize a new PlannerTask object
    viper, err := viperConfig.InitViper("config.json")
	requestBody := models.NewPlannerTask()

	// Retrieve planId and bucketId from viper configuration
  
	planId := viper.GetString("planId")
	bucketId := viper.GetString("bucketId")

	// Set the planId, bucketId, and title for the task
	requestBody.SetPlanId(&planId)
	requestBody.SetBucketId(&bucketId)
	requestBody.SetTitle(&taskTitle) // Changed to use function parameter
	result, err := g.appClient.Planner().Tasks().Post(context.Background(), requestBody, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GraphHelper) GetTeamById(teamId string) (models.Teamable, error) {
  team, nil := g.appClient.Teams().ByTeamId(teamId).Get(context.Background(), nil)
  return team, nil
}

func (g *GraphHelper) GetAllChannels(teamId string) (models.ChannelCollectionResponseable, error) {
  channels, nil := g.appClient.Teams().ByTeamId(teamId).AllChannels().Get(context.Background(), nil)
  return channels, nil
}

func (g *GraphHelper) GetChannelById(teamId string, channelId string) (models.Channelable, error) {
    channel, nil := g.appClient.Teams().ByTeamId(teamId).Channels().ByChannelId(channelId).Get(context.Background(), nil)
    return channel, nil
    }
    
    func (g *GraphHelper) EnsureFilesFolder(teamId string, channelId string) (models.DriveItemable, error) {
    drive, nil := g.appClient.Teams().ByTeamId(teamId).Channels().ByChannelId(channelId).FilesFolder().Get(context.Background(), nil)
    return drive, nil
    }
        func (g *GraphHelper) GetTabs(teamId string, channelId string) (models.TeamsTabCollectionResponseable, error) {
    teamTabs, nil := g.appClient.Teams().ByTeamId(teamId).Channels().ByChannelId(channelId).Tabs().Get(context.Background(), nil)
    return teamTabs, nil
    }