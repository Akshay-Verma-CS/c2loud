package gcp

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	appengine "google.golang.org/api/appengine/v1"
	cloudfunctions "google.golang.org/api/cloudfunctions/v1"
	compute "google.golang.org/api/compute/v1"
	container "google.golang.org/api/container/v1"
)

// GCPConfig defines the configuration for GCPProvider.
type GCPConfig struct {
	UseIAMRole  bool
	Credentials string // Path to JSON key file
	ProjectID   string
	Region      string
}

// GCPProvider holds the clients for various GCP services.
type GCPProvider struct {
	ComputeService          *compute.Service
	AppEngineService        *appengine.Service
	KubernetesEngineService *container.Service
	CloudFunctionsService   *cloudfunctions.Service
	// ... add other service clients as needed ...
}

var (
	instances = make(map[string]*GCPProvider) // Holds the map of GCPProvider instances
	mutex     sync.Mutex                      // Ensures thread-safe access to the instances map
)

// NewGCPProvider creates a new GCPProvider with all the necessary service clients.
func NewGCPProvider(ctx context.Context, config *GCPConfig) (*GCPProvider, error) {
	var clientOption google.ClientOption

	if config.UseIAMRole {
		if metadata.OnGCE() {
			// Use the App Engine default service account
			clientOption = google.ComputeTokenSource("")
		} else {
			return nil, fmt.Errorf("UseIAMRole is true, but not running on GCE or App Engine")
		}
	} else {
		creds, err := google.CredentialsFromJSON(ctx, []byte(config.Credentials), compute.ComputeScope)
		if err != nil {
			return nil, err
		}
		clientOption = google.WithCredentials(creds)
	}

	computeService, err := compute.NewService(ctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Compute service: %v", err)
	}

	appengineService, err := appengine.NewService(ctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("Failed to create App Engine service: %v", err)
	}

	kubernetesService, err := container.NewService(ctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kubernetes Engine service: %v", err)
	}

	cloudfunctionsService, err := cloudfunctions.NewService(ctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Cloud Functions service: %v", err)
	}

	return &GCPProvider{
		ComputeService:          computeService,
		AppEngineService:        appengineService,
		KubernetesEngineService: kubernetesService,
		CloudFunctionsService:   cloudfunctionsService,
		// ... initialize other services ...
	}, nil
}

func GetInstance(ctx context.Context, key string, config *GCPConfig) (*GCPProvider, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if instance, exists := instances[key]; exists {
		// Return existing instance
		return instance, nil
	}

	// Create a new instance if it does not exist
	instance, err := NewGCPProvider(ctx, config)
	if err != nil {
		return nil, err
	}

	instances[key] = instance
	return instance, nil
}
