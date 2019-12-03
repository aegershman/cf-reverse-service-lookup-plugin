package apihelper

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/cfcurl"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/models"
)

// CFAPIHelper wraps cf curl results
type CFAPIHelper interface {
	GetServiceInstance(string) (models.Service, error)
}

// APIHelper -
type APIHelper struct {
	cli plugin.CliConnection
}

// New -
func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}

// GetServiceInstance -
func (api *APIHelper) GetServiceInstance(serviceGUID string) (models.Service, error) {
	summaryJSON, err := cfcurl.Curl(api.cli, "service_instances/"+serviceGUID)
	if err != nil {
		return models.Service{}, err
	}

	entity := summaryJSON["entity"].(map[string]interface{})

	return models.Service{
		Name:     entity["name"].(string),
		SpaceURL: entity["space_url"].(string),
	}, nil
}
