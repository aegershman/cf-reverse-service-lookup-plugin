package apihelper

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/models"
)

// CFAPIHelper wraps cf curl results
type CFAPIHelper interface {
	GetService(string) (models.Service, error)
}

// APIHelper -
type APIHelper struct {
	cli plugin.CliConnection
}

// New -
func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}

// GetService -
func (api *APIHelper) GetService(serviceGUID string) (models.Service, error) {
	return models.Service{}, nil
}
