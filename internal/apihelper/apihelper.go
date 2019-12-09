package apihelper

import (
	"errors"
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/internal/cfcurl"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/internal/models"
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

	if entity, ok := summaryJSON["entity"].(map[string]interface{}); ok {
		return models.Service{
			Name:      entity["name"].(string),
			SpaceGUID: entity["space_guid"].(string),
		}, nil
	}

	return models.Service{}, errors.New(fmt.Sprintln(summaryJSON))

}

// GetSpaceByGUID -
func (api *APIHelper) GetSpaceByGUID(spaceGUID string) (models.Space, error) {
	path := fmt.Sprintf("/v2/spaces/%s", spaceGUID)
	summaryJSON, err := cfcurl.Curl(api.cli, path)
	if err != nil {
		return models.Space{}, err
	}
	log.Traceln(summaryJSON)

	if entity, ok := summaryJSON["entity"].(map[string]interface{}); ok {
		return models.Space{
			Name:             entity["name"].(string),
			OrganizationGUID: entity["organization_guid"].(string),
		}, nil
	}

	return models.Space{}, errors.New(fmt.Sprintln(summaryJSON))

}

// GetOrganizationByGUID -
func (api *APIHelper) GetOrganizationByGUID(organizationGUID string) (models.Organization, error) {
	path := fmt.Sprintf("/v2/organizations/%s", organizationGUID)
	summaryJSON, err := cfcurl.Curl(api.cli, path)
	if err != nil {
		return models.Organization{}, err
	}
	log.Traceln(summaryJSON)

	if entity, ok := summaryJSON["entity"].(map[string]interface{}); ok {
		return models.Organization{
			Name: entity["name"].(string),
		}, nil
	}

	return models.Organization{}, errors.New(fmt.Sprintln(summaryJSON))

}
