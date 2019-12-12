package v2client

// Organization -
type Organization struct {
	Name string `json:"name"`
}

// OrgsService -
type OrgsService service

// GetOrganizationByGUID -
func (o *OrgsService) GetOrganizationByGUID(orgGUID string) (Organization, error) {
	org, err := o.client.cfc.GetOrgByGuid(orgGUID)
	if err != nil {
		return Organization{}, err
	}

	return Organization{
		Name: org.Name,
	}, nil
}
