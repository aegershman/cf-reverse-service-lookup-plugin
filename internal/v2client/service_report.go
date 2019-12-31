package v2client

// ServiceReport -
type ServiceReport struct {
	Organization `json:"organization"`
	Service      `json:"service"`
	Space        `json:"space"`
}

// ServiceReportService -
type ServiceReportService service

// GetServiceReportsFromServiceGUIDs -
func (s *ServiceReportService) GetServiceReportsFromServiceGUIDs(serviceGUIDs ...string) ([]ServiceReport, error) {
	var serviceReports []ServiceReport
	for _, serviceGUID := range serviceGUIDs {
		serviceReport, err := s.client.ServiceReportService.GetServiceReportFromServiceGUID(serviceGUID)
		if err != nil {
			return nil, err
		}
		serviceReports = append(serviceReports, serviceReport)
	}
	return serviceReports, nil
}

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
