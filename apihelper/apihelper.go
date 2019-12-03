package apihelper

import (
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/cfcurl"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/models"
	log "github.com/sirupsen/logrus"
)

// CFAPIHelper wraps cf-curl results, acts as a cf-curl client
type CFAPIHelper interface {
	GetServiceInstanceByGUID(string) (models.Service, error)
	GetSpaceByGUID(string) (models.Space, error)
	GetOrganizationByGUID(string) (models.Organization, error)
}

// APIHelper -
type APIHelper struct {
	cli plugin.CliConnection
}

// New -
func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}

// GetServiceInstanceByGUID -
func (api *APIHelper) GetServiceInstanceByGUID(serviceGUID string) (models.Service, error) {
	path := fmt.Sprintf("/v2/service_instances/%s", serviceGUID)
	summaryJSON, err := cfcurl.Curl(api.cli, path)
	if err != nil {
		return models.Service{}, err
	}
	log.Traceln(summaryJSON)

	entity := summaryJSON["entity"].(map[string]interface{})

	return models.Service{
		Name:      entity["name"].(string),
		SpaceGUID: entity["space_guid"].(string),
	}, nil
}

// GetSpaceByGUID -
func (api *APIHelper) GetSpaceByGUID(spaceGUID string) (models.Space, error) {
	path := fmt.Sprintf("/v2/spaces/%s", spaceGUID)
	summaryJSON, err := cfcurl.Curl(api.cli, path)
	if err != nil {
		return models.Space{}, err
	}
	log.Traceln(summaryJSON)

	entity := summaryJSON["entity"].(map[string]interface{})

	return models.Space{
		Name:             entity["name"].(string),
		OrganizationGUID: entity["organization_guid"].(string),
	}, nil
}

// GetOrganizationByGUID -
func (api *APIHelper) GetOrganizationByGUID(organizationGUID string) (models.Organization, error) {
	path := fmt.Sprintf("/v2/organizations/%s", organizationGUID)
	summaryJSON, err := cfcurl.Curl(api.cli, path)
	if err != nil {
		return models.Organization{}, err
	}
	log.Traceln(summaryJSON)

	entity := summaryJSON["entity"].(map[string]interface{})

	return models.Organization{
		Name: entity["name"].(string),
	}, nil
}
