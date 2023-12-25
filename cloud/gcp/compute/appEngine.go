package compute

import (
	"context"
	"fmt"

	"google.golang.org/api/appengine/v1"
)

// AppEngineService provides operations for interacting with GCP App Engine.
type AppEngineService struct {
	service *appengine.Service
	ctx     context.Context
}

// NewAppEngineService creates a new instance of AppEngineService.
func NewAppEngineService(ctx context.Context) (*AppEngineService, error) {
	service, err := appengine.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &AppEngineService{
		service: service,
		ctx:     ctx,
	}, nil
}

// CreateApplication creates a new App Engine application in a given location.
func (ae *AppEngineService) CreateApplication(locationID, projectID string) (*appengine.Operation, error) {
	// Define the application to be created
	application := &appengine.Application{
		Id:         projectID,  // Set the Project ID
		LocationId: locationID, // Set the location, e.g., "us-central"
	}

	// Call App Engine's create application method
	op, err := ae.service.Apps.Create(application).Context(ae.ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("App Engine Application creation failed: %v", err)
	}
	return op, nil
}

// ListApplications lists all App Engine applications in a given project.
func (ae *AppEngineService) ListApplications() ([]*appengine.Application, error) {
	appsListCall := ae.service.Apps.List()
	response, err := appsListCall.Context(ae.ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("App Engine Application listing failed: %v", err)
	}
	return response.Apps, nil
}

// GetApplication retrieves details of a specific App Engine application.
func (ae *AppEngineService) GetApplication(appID string) (*appengine.Application, error) {
	app, err := ae.service.Apps.Get(appID).Context(ae.ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error retrieving App Engine application: %v", err)
	}
	return app, nil
}

// UpdateApplication updates specific fields of an App Engine application.
func (ae *AppEngineService) UpdateApplication(appID string, updates *appengine.Application) (*appengine.Operation, error) {
	op, err := ae.service.Apps.Patch(appID, updates).Context(ae.ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error updating App Engine application: %v", err)
	}
	return op, nil
}
