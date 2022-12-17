package v2client

import "github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/details"

func (c *Client) GetOrganizationByGUID(orgGUID string) (details.Organization, error) {
	org, err := c.cfc.GetOrgByGuid(orgGUID)
	if err != nil {
		return details.Organization{}, err
	}

	organization := details.Organization{
		Name: org.Name,
	}
	return organization, nil
}
