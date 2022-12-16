package v2client

import "github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/details"

func (c *Client) GetSpaceByGUID(spaceGUID string) (details.Space, error) {
	space, err := c.cfc.GetSpaceByGuid(spaceGUID)
	if err != nil {
		return details.Space{}, err
	}

	d := details.Space{
		Name:             space.Name,
		OrganizationGUID: space.OrganizationGuid,
	}
	return d, nil
}
