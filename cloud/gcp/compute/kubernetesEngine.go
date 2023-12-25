package compute

import (
	"google.golang.org/api/container/v1"
)

type KubernetesEngineService struct {
	service *container.Service
}

func NewKubernetesEngineService() *KubernetesEngineService {
	// Initialize Kubernetes Engine Service
	containerService, err := container.New(client)
	if err != nil {
		// handle error
	}

	return &KubernetesEngineService{
		service: containerService,
	}
}

// CreateCluster creates a new Kubernetes cluster.
func (ke *KubernetesEngineService) CreateCluster(projectID, zone string, cluster *container.Cluster) error {
	_, err := ke.service.Projects.Zones.Clusters.Create(projectID, zone, cluster).Do()
	return err
}

// ListClusters lists all clusters in a given zone.
func (ke *KubernetesEngineService) ListClusters(projectID, zone string) ([]*container.Cluster, error) {
	response, err := ke.service.Projects.Zones.Clusters.List(projectID, zone).Do()
	if err != nil {
		return nil, err
	}
	return response.Clusters, nil
}
