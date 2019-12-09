package v2client

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

// Client -
type Client struct {
	cli    plugin.CliConnection
	common service

	Orgs     *OrgsService
	Services *ServicesService
	Spaces   *SpacesService
}

// NewClient -
func NewClient(cli plugin.CliConnection) *Client {
	c := &Client{cli: cli}
	c.common.client = c
	c.Orgs = (*OrgsService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.Spaces = (*SpacesService)(&c.common)
	return c
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
