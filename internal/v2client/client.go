package v2client

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/cloudfoundry/cli/plugin"
)

// Client -
type Client struct {
	cfc    cfclient.CloudFoundryClient
	cli    plugin.CliConnection
	common service

	Orgs     *OrgsService
	Services *ServicesService
	Spaces   *SpacesService
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

	c := &Client{cli: cli}
	c.cfc = cfc
	c.common.client = c
	c.Orgs = (*OrgsService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.Spaces = (*SpacesService)(&c.common)
	return c, nil
}

type service struct {
	client *Client
}

// Curl -
func (c *Client) Curl(path string) (map[string]interface{}, error) {
	output, err := c.cli.CliCommandWithoutTerminalOutput("curl", path)
	if err != nil {
		return nil, err
	}

	if nil == output || 0 == len(output) {
		return nil, errors.New("CF API returned no output")
	}

	data := strings.Join(output, "\n")
	if 0 == len(data) || "" == data {
		return nil, errors.New("Failed to join output")
	}

	var f interface{}
	err = json.Unmarshal([]byte(data), &f)
	return f.(map[string]interface{}), err
}
