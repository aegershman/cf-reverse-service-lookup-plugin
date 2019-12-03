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
	path := fmt.Sprintf("/v2/service_instances/%s", serviceGUID)
	summaryJSON, err := cfcurl.Curl(api.cli, path)
	if err != nil {
		return models.Service{}, err
	}
	log.Traceln(summaryJSON)

	entity := summaryJSON["entity"].(map[string]interface{})

	return models.Service{
		Name:     entity["name"].(string),
		SpaceURL: entity["space_url"].(string),
	}, nil
}
