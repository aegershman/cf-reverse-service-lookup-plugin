package v2client

import (
	"errors"
	"fmt"
)

// Space -
type Space struct {
	Name             string `json:"name"`
	OrganizationGUID string `json:"organization_guid"`
}

// SpacesService -
type SpacesService service

// GetSpaceByGUID -
func (s *SpacesService) GetSpaceByGUID(spaceGUID string) (Space, error) {
	path := fmt.Sprintf("/v2/spaces/%s", spaceGUID)
	summaryJSON, err := s.client.Curl(path)
	if err != nil {
		return Space{}, err
	}

	if entity, ok := summaryJSON["entity"].(map[string]interface{}); ok {
		return Space{
			Name:             entity["name"].(string),
			OrganizationGUID: entity["organization_guid"].(string),
		}, nil
	}

	return Space{}, errors.New(fmt.Sprintln(summaryJSON))

}
