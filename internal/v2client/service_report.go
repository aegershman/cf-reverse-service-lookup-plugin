package v2client

// ServiceReport -
type ServiceReport struct {
	Organization `json:"organization"`
	Service      `json:"service"`
	Space        `json:"space"`
}

// ServiceReportService -
type ServiceReportService service

// GetServiceReportFromServiceGUID -
func (s *ServiceReportService) GetServiceReportFromServiceGUID(serviceGUID string) (ServiceReport, error) {
	serviceInstance, err := s.client.Services.GetServiceInstanceByGUID(serviceGUID)
	if err != nil {
		return ServiceReport{}, err
	}

	serviceSpace, err := s.client.Spaces.GetSpaceByGUID(serviceInstance.SpaceGUID)
	if err != nil {
		return ServiceReport{}, err
	}

	serviceOrganization, err := s.client.Orgs.GetOrganizationByGUID(serviceSpace.OrganizationGUID)
	if err != nil {
		return ServiceReport{}, err
	}

	return ServiceReport{
		Service:      serviceInstance,
		Space:        serviceSpace,
		Organization: serviceOrganization,
	}, nil

}
