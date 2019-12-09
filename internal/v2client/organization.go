package v2client

import (
	"errors"
	"fmt"
)

// Organization -
type Organization struct {
	Name string `json:"name"`
}

// OrgsService -
type OrgsService service

// GetOrganizationByGUID -
func (o *OrgsService) GetOrganizationByGUID(organizationGUID string) (Organization, error) {
	path := fmt.Sprintf("/v2/organizations/%s", organizationGUID)
	summaryJSON, err := o.client.Curl(path)
	if err != nil {
		return Organization{}, err
	}

	if entity, ok := summaryJSON["entity"].(map[string]interface{}); ok {
		return Organization{
			Name: entity["name"].(string),
		}, nil
	}

	return Organization{}, errors.New(fmt.Sprintln(summaryJSON))

}
