package compute

import (
	"google.golang.org/api/cloudfunctions/v1"
)

type CloudFunctionsService struct {
	service *cloudfunctions.Service
}

func NewCloudFunctionsService() *CloudFunctionsService {
	// Initialize Cloud Functions Service
	cloudFunctionsService, err := cloudfunctions.New(client)
	if err != nil {
		// handle error
	}

	return &CloudFunctionsService{
		service: cloudFunctionsService,
	}
}

// CreateFunction creates a new cloud function.
func (cf *CloudFunctionsService) CreateFunction(projectLocation string, function *cloudfunctions.CloudFunction) error {
	_, err := cf.service.Projects.Locations.Functions.Create(projectLocation, function).Do()
	return err
}

// ListFunctions lists all cloud functions in a given location.
func (cf *CloudFunctionsService) ListFunctions(projectLocation string) ([]*cloudfunctions.CloudFunction, error) {
	response, err := cf.service.Projects.Locations.Functions.List(projectLocation).Do()
	if err != nil {
		return nil, err
	}
	return response.Functions, nil
}
