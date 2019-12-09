package v2client

import (
	"errors"
	"fmt"
)

// Service -
type Service struct {
	Name      string `json:"name"`
	SpaceGUID string `json:"space_guid"`
}

// ServicesService -
type ServicesService service

// GetServiceInstanceByGUID -
func (s *ServicesService) GetServiceInstanceByGUID(serviceGUID string) (Service, error) {
	path := fmt.Sprintf("/v2/service_instances/%s", serviceGUID)
	summaryJSON, err := s.client.Curl(path)
	if err != nil {
		return Service{}, err
	}

	if entity, ok := summaryJSON["entity"].(map[string]interface{}); ok {
		return Service{
			Name:      entity["name"].(string),
			SpaceGUID: entity["space_guid"].(string),
		}, nil
	}

	return Service{}, errors.New(fmt.Sprintln(summaryJSON))

}
