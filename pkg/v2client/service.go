package v2client

import "github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/details"

func (c *Client) GetServiceInstanceByGUID(serviceGUID string) (details.Service, error) {
	service, err := c.cfc.ServiceInstanceByGuid(serviceGUID)
	if err != nil {
		return details.Service{}, err
	}

	s := details.Service{
		GUID:      service.Guid,
		Name:      service.Name,
		SpaceGUID: service.SpaceGuid,
	}
	return s, nil
}
