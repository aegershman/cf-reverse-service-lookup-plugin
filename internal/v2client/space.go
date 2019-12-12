package v2client

// Space -
type Space struct {
	Name             string `json:"name"`
	OrganizationGUID string `json:"organization_guid"`
}

// SpacesService -
type SpacesService service

// GetSpaceByGUID -
func (s *SpacesService) GetSpaceByGUID(spaceGUID string) (Space, error) {
	space, err := s.client.cfc.GetSpaceByGuid(spaceGUID)
	if err != nil {
		return Space{}, err
	}

	return Space{
		Name:             space.Name,
		OrganizationGUID: space.OrganizationGuid,
	}, nil
}
