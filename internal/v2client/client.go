package v2client

import (
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/cloudfoundry-community/go-cfclient"
)

type service struct {
	client *Client
}

// Client -
type Client struct {
	cfc    cfclient.CloudFoundryClient
	common service

	Orgs                 *OrgsService
	ServiceReportService *ServiceReportService
	Services             *ServicesService
	Spaces               *SpacesService
}

// NewClient -
func NewClient(cli plugin.CliConnection) (*Client, error) {
	apiAddress, err := cli.ApiEndpoint()
	if err != nil {
		return &Client{}, nil
	}

	accessToken, err := cli.AccessToken()
	if err != nil {
		return &Client{}, nil
	}

	trimmedAccessToken := strings.TrimPrefix(accessToken, "bearer ")

	cfcConfig := &cfclient.Config{
		ApiAddress: apiAddress,
		Token:      trimmedAccessToken,
	}

	cfc, err := cfclient.NewClient(cfcConfig)
	if err != nil {
		return &Client{}, nil
	}

	c := &Client{}
	c.cfc = cfc
	c.common.client = c
	c.Orgs = (*OrgsService)(&c.common)
	c.ServiceReportService = (*ServiceReportService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.Spaces = (*SpacesService)(&c.common)
	return c, nil
}
