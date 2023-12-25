package compute

import (
	"google.golang.org/api/compute/v1"
)

type ComputeEngineService struct {
	service *compute.Service
}

func NewComputeEngineService() *ComputeEngineService {
	// Assume 'client' is the authenticated client for GCP
	computeService, err := compute.New(client)
	if err != nil {
		// handle error
	}

	return &ComputeEngineService{
		service: computeService,
	}
}

// AddVMInstance creates a new VM instance.
func (ce *ComputeEngineService) AddVMInstance(projectID, zone string, instance *compute.Instance) error {
	_, err := ce.service.Instances.Insert(projectID, zone, instance).Do()
	return err
}

// ListVMInstances lists all VM instances in a given zone.
func (ce *ComputeEngineService) ListVMInstances(projectID, zone string) ([]*compute.Instance, error) {
	response, err := ce.service.Instances.List(projectID, zone).Do()
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}
