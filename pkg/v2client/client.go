package v2client

import (
	"github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/details"
	"github.com/cloudfoundry-community/go-cfclient"
)

type Client struct {
	cfc cfclient.CloudFoundryClient
}

func NewClient(apiAddress string, accessToken string) (*Client, error) {
	cfcConfig := &cfclient.Config{
		ApiAddress: apiAddress,
		Token:      accessToken,
	}

	cfc, err := cfclient.NewClient(cfcConfig)
	if err != nil {
		return &Client{}, err
	}

	r := &Client{
		cfc: cfc,
	}
	return r, nil
}

func (c *Client) GetServiceInstancesDetails(serviceGUIDs ...string) ([]details.ServiceInstanceDetails, error) {
	var serviceInstanceDetails []details.ServiceInstanceDetails
	for _, guid := range serviceGUIDs {
		summary, err := c.GetServiceInstanceDetails(guid)
		if err != nil {
			return nil, err
		}
		serviceInstanceDetails = append(serviceInstanceDetails, summary)
	}
	return serviceInstanceDetails, nil
}

func (c *Client) GetServiceInstanceDetails(serviceGUID string) (details.ServiceInstanceDetails, error) {
	service, err := c.GetServiceInstanceByGUID(serviceGUID)
	if err != nil {
		return details.ServiceInstanceDetails{}, err
	}

	space, err := c.GetSpaceByGUID(service.SpaceGUID)
	if err != nil {
		return details.ServiceInstanceDetails{}, err
	}

	org, err := c.GetOrganizationByGUID(space.OrganizationGUID)
	if err != nil {
		return details.ServiceInstanceDetails{}, err
	}

	d := details.ServiceInstanceDetails{
		Organization: org,
		Space:        space,
		Service:      service,
	}
	return d, nil
}
