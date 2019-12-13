package v2client

// Service -
type Service struct {
	GUID      string `json:"guid"`
	Name      string `json:"name"`
	SpaceGUID string `json:"space_guid"`
}

// ServicesService -
type ServicesService service

// GetServiceInstanceByGUID -
func (s *ServicesService) GetServiceInstanceByGUID(serviceGUID string) (Service, error) {
	service, err := s.client.cfc.ServiceInstanceByGuid(serviceGUID)
	if err != nil {
		return Service{}, err
	}

	return Service{
		GUID:      service.Guid,
		Name:      service.Name,
		SpaceGUID: service.SpaceGuid,
	}, nil
}
